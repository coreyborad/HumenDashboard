package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"stock/models"
	"stock/repositories"
	"stock/websocket"
	"strconv"
	"strings"
	"sync"
	"time"
)

// StockService Stock Service
type StockService struct {
	stockRep *repositories.StockRepository
}

// NewStockService New Stock Service
func NewStockService(
	stockRep *repositories.StockRepository,
) *StockService {
	return &StockService{
		stockRep: stockRep,
	}
}

func (s *StockService) ReadWs(ctx context.Context, client *websocket.Client) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-client.ReadMessage:
			if !ok {
				break
			}
			var data struct {
				Type string          `json:"type"`
				Data json.RawMessage `json:"data"`
			}
			err := json.Unmarshal(msg, &data)
			if err != nil {
				break
			}
			// 先了解是傳甚麼資料過來
			if data.Type == "stock" {
				filter := models.WsStockData{}
				err := json.Unmarshal(data.Data, &filter)
				if err != nil {
					fmt.Println(err)
					break
				}
				// 檢查欄位
				if filter.StockNumber == "" {
					break
				}
				if filter.StartDate == nil {
					break
				}
				if filter.EndDate == nil {
					break
				}
				to_client := &models.WsToClientData{}
				stock, err := s.stockRep.GetStockDataByDate(filter.StockNumber, filter.StartDate, filter.EndDate)
				to_client.Data = stock
				jsondata, err := json.Marshal(to_client)
				client.Send(jsondata)
			}
		}
	}
}

func (s *StockService) DailyParser() error {
	// Get data
	url := "https://www.twse.com.tw/exchangeReport/STOCK_DAY_ALL?response=json"
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	type RawPayload struct {
		Data [][]string `json:"data"`
		Date string     `json:"date"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rawpayload := &RawPayload{}
	json.Unmarshal(body, rawpayload)
	// Format Data with go routine
	dealDateFormat, err := time.Parse("20060102", rawpayload.Date)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dealDate := time.Date(dealDateFormat.Year(), dealDateFormat.Month(), dealDateFormat.Day(), 8, 0, 0, 0, time.Now().Location())
	replaceTenPercentile := func(target string) string {
		return strings.ReplaceAll(target, ",", "")
	}
	wg := sync.WaitGroup{}
	for _, stock := range rawpayload.Data {
		now := time.Now()
		stockName := stock[1]
		dealCount, _ := strconv.ParseInt(replaceTenPercentile(stock[2]), 10, 64)
		priceOnOpen, _ := strconv.ParseFloat(replaceTenPercentile(stock[4]), 64)
		priceOnHighest, _ := strconv.ParseFloat(replaceTenPercentile(stock[5]), 64)
		priceOnLowest, _ := strconv.ParseFloat(replaceTenPercentile(stock[6]), 64)
		priceOnClose, _ := strconv.ParseFloat(replaceTenPercentile(stock[7]), 64)
		stockData := &models.StockData{
			StockNumber:    stock[0],
			DealDate:       &dealDate,
			DealCount:      uint64(dealCount),
			PriceOnOpen:    priceOnOpen,
			PriceOnHighest: priceOnHighest,
			PriceOnLowest:  priceOnLowest,
			PriceOnClose:   priceOnClose,
			CreatedAt:      &now,
			UpdatedAt:      &now,
		}
		_, err := s.stockRep.Find(stockData.StockNumber)
		if err != nil {
			fmt.Println("empty")
			stockModel := &models.Stock{}
			stockModel.StockNumber = stockData.StockNumber
			stockModel.StockName = stockName
			stockModel.CreatedAt = stockData.CreatedAt
			stockModel.UpdatedAt = stockData.UpdatedAt
			s.stockRep.Create(stockModel)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = s.stockRep.InsertToMongo(stockData)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
	return nil
}

func (s *StockService) Calc(stockNumberList []string, date *time.Time) (interface{}, error) {
	if stockNumberList == nil || len(stockNumberList) <= 0 {
		stockNumberList = []string{
			"0050",
			"2330",
			"00733",
			"2412",
		}
	}
	if date == nil {
		now := time.Now()
		date = &now
	}
	for _, stockNumber := range stockNumberList {
		fmt.Println("====Start====")
		fmt.Printf("%s => %s  \n", date.String(), stockNumber)
		lastDays := int64(60)
		// Desc order
		stockList, err := s.stockRep.GetLastStockData("0050", &lastDays, date)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if len(stockList) <= 0 {
			fmt.Println("Empty skip calc")
			return nil, nil
		}

		// ----Calc KD----
		// 策略
		// 1. 碰到Uppercross = true 隔日做多
		// 直到K跟D都超過80就賣出
		// 2. 碰到High Lag = true 隔日做多
		// 2天後賣出
		lastKD := int64(9)
		// ASC order
		lastKDList, err := s.stockRep.CalcKDVal(stockList[:lastKD], &lastKD)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		// for _, v := range lastKDList {
		// 	fmt.Println(v.KVal, v.DVal)
		// }
		// 0. 標示出最後一筆K跟D
		fmt.Printf("K: %f, D: %f \n", lastKDList[int(lastKD)-1].KVal, lastKDList[int(lastKD)-1].DVal)
		// 1. KD上漲交叉
		// Condtion(n=9)
		// 012345678
		// [2:7] => 23456
		// A: 最後1筆KD要小於80
		// B: 第2~第6筆,要都是D>K
		// C: 最後2筆KD值,要都是K>D
		checkUppercrossFunc := func() bool {
			if lastKDList[int(lastKD)-1].KVal > 80 || lastKDList[int(lastKD)-1].DVal > 80 {
				return false
			}
			for _, kdVal := range lastKDList[2 : (int(lastKD))-2] {
				if kdVal.DVal < kdVal.KVal {
					return false
				}
			}
			for _, kdVal := range lastKDList[(int(lastKD))-2:] {
				if kdVal.KVal < kdVal.DVal {
					return false
				}
			}
			return true
		}
		isUppercross := checkUppercrossFunc()
		fmt.Printf("Uppercross: %t \n", isUppercross)
		// 2. KD下跌交叉
		// Condtion(n=9)
		// A: 最後1筆KD要大於20
		// B: 第2~第6筆,要都是D<K
		// C: 最後2筆KD值,要都是K<D
		checkUndercrossFunc := func() bool {
			if lastKDList[int(lastKD)-1].KVal < 20 || lastKDList[int(lastKD)-1].DVal < 20 {
				return false
			}
			for _, kdVal := range lastKDList[2 : (int(lastKD))-2] {
				if kdVal.DVal > kdVal.KVal {
					return false
				}
			}
			for _, kdVal := range lastKDList[(int(lastKD))-2:] {
				if kdVal.KVal > kdVal.DVal {
					return false
				}
			}
			return true
		}
		isUndercross := checkUndercrossFunc()
		fmt.Printf("Undercross: %t \n", isUndercross)
		// 3. KD高檔鈍化
		// Condtion(n=9)
		// A: 最後3筆K>80
		// B: 最後1筆收盤價>最後第2,3筆的收盤價
		// C: 最後1筆K > D
		checkHighLag := func() bool {
			for _, kdVal := range lastKDList[(int(lastKD))-3:] {
				if kdVal.KVal < 80 {
					return false
				}
			}
			lastClosePrice := stockList[0].PriceOnClose
			if lastClosePrice < stockList[1].PriceOnClose || lastClosePrice < stockList[2].PriceOnClose {
				return false
			}
			if lastKDList[int(lastKD)-1].KVal < lastKDList[int(lastKD)-1].DVal {
				return false
			}
			return true
		}
		isHighLag := checkHighLag()
		fmt.Printf("High Lag: %t \n", isHighLag)
		// 4. KD低檔鈍化
		// Condtion(n=9)
		// A: 最後3筆K<20
		// B: 最後1筆收盤價<最後第2,3筆的收盤價
		// C: 最後1筆K < D
		checkLowLag := func() bool {
			for _, kdVal := range lastKDList[(int(lastKD))-3:] {
				if kdVal.KVal > 20 {
					return false
				}
			}
			lastClosePrice := stockList[0].PriceOnClose
			if lastClosePrice > stockList[1].PriceOnClose || lastClosePrice > stockList[2].PriceOnClose {
				return false
			}
			if lastKDList[int(lastKD)-1].KVal > lastKDList[int(lastKD)-1].DVal {
				return false
			}
			return true
		}
		isLowLag := checkLowLag()
		fmt.Printf("Low Lag: %t \n", isLowLag)
	}

	return nil, nil
}

func (s *StockService) ParserDataOnManual() error {
	now := time.Now()
	s.stockRep.ParserStockDataByMonth("0050", &now)
	return nil
}
