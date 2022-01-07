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
		wg.Add(1)
		go func() {
			defer wg.Done()
			stockModel, err := s.stockRep.Find(stockData.StockNumber)
			if err != nil {
				fmt.Println("empty")
				stockModel.StockNumber = stockData.StockNumber
				stockModel.StockName = stockName
				stockModel.CreatedAt = stockData.CreatedAt
				stockModel.UpdatedAt = stockData.UpdatedAt
				s.stockRep.Create(stockModel)
			}
			err = s.stockRep.InsertToMongo(stockData)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
	return nil
}

func (s *StockService) DailyCalc() error {
	fmt.Println("====Start====")
	last := int64(18)

	stockList, err := s.stockRep.GetLastStockData("0050", &last)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if len(stockList) <= 0 {
		fmt.Println("Empty skip calc")
		return nil
	}
	periodHighest := stockList[len(stockList)-1].PriceOnHighest
	periodLowest := stockList[len(stockList)-1].PriceOnLowest
	lastKVal := float64(50)
	lastDVal := float64(50)
	for i := len(stockList) - 1; i >= 0; i-- {
		if stockList[i].PriceOnHighest > periodHighest {
			periodHighest = stockList[i].PriceOnHighest
		}
		if stockList[i].PriceOnLowest < periodLowest {
			periodLowest = stockList[i].PriceOnLowest
		}
		closePrice := stockList[i].PriceOnClose

		rsv := (closePrice - periodLowest) / (periodHighest - periodLowest) * 100
		lastKVal = (0.67 * lastKVal) + (0.33 * rsv)
		lastDVal = (0.67 * lastDVal) + (0.33 * lastKVal)
	}
	fmt.Println(lastKVal, lastDVal)
	return nil
}

func (s *StockService) ParserDataOnManual() error {
	now := time.Now()
	s.stockRep.ParserStockDataByMonth("0050", &now)
	return nil
}
