package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"stock/models"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	err = errors.New("Error")
)

// StockRepository StockRepository
type StockRepository struct {
	db      *gorm.DB
	mongodb *mongo.Collection
}

// StockRepository New StockHmi StockRepository
func NewStockRepository(db *gorm.DB, mongodb *mongo.Database) *StockRepository {
	return &StockRepository{
		db:      db,
		mongodb: mongodb.Collection("stock_history"),
	}
}

// Find Find
func (r *StockRepository) Find(stockNumber string) (*models.Stock, error) {
	stock := new(models.Stock)
	if err := r.db.Where("`stock_number` = ?", stockNumber).First(&stock).Error; err != nil {
		return nil, err
	}

	return stock, nil
}

func (r *StockRepository) Create(stock *models.Stock) (*models.Stock, error) {
	if err := r.db.Create(&stock).Error; err != nil {
		return nil, err
	}

	return stock, nil
}

// Find Find
func (r *StockRepository) GetStockDataByDate(stockNumber string, start_date *time.Time, end_date *time.Time) ([]*models.StockData, error) {
	ctx := context.Background()
	filter := bson.M{
		"stock_number": stockNumber,
		"deal_date": bson.M{
			"$gt": start_date.UTC(),
			"$lt": end_date.UTC(),
		},
	}
	cursor, err := r.mongodb.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	stock_list := []*models.StockData{}
	if err = cursor.All(ctx, &stock_list); err != nil {
		return nil, err
	}
	return stock_list, nil
}

func (r *StockRepository) GetLastStockData(stockNumber string, last *int64, beforeThisDate *time.Time) ([]*models.StockData, error) {
	ctx := context.Background()
	fmt.Println(beforeThisDate.UTC())
	filter := bson.M{
		"stock_number": stockNumber,
		"deal_date": bson.M{
			"$lte": beforeThisDate.UTC(),
		},
	}

	cursor, err := r.mongodb.Find(ctx, filter, &options.FindOptions{
		Sort: map[string]interface{}{
			"deal_date": -1,
		},
		Limit: last,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	stock_list := []*models.StockData{}
	if err = cursor.All(ctx, &stock_list); err != nil {
		return nil, err
	}
	// Desc order
	return stock_list, nil
}

func (r *StockRepository) ParserStockDataByMonth(stockNumber string, date *time.Time) error {
	ctx := context.Background()
	// Get data
	url := fmt.Sprintf("https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=%s&stockNo=%s", date.Format("201601"), stockNumber)
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	type RawPayload struct {
		Data [][]string `json:"data"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rawpayload := &RawPayload{}
	json.Unmarshal(body, rawpayload)
	// Format data
	replaceTenPercentile := func(target string) string {
		return strings.ReplaceAll(target, ",", "")
	}
	stock := &models.StockData{}
	for _, dailyData := range rawpayload.Data {
		formatDate := strings.Split(dailyData[0], "/")
		now := time.Now()
		year, _ := strconv.Atoi(formatDate[0])
		month, _ := strconv.Atoi(formatDate[1])
		day, _ := strconv.Atoi(formatDate[2])
		// We are UTC+8
		dealDate := time.Date(year+1911, time.Month(month), day, 8, 0, 0, 0, time.Now().Location())
		dealCount, _ := strconv.ParseInt(replaceTenPercentile(dailyData[1]), 10, 64)
		priceOnOpen, _ := strconv.ParseFloat(replaceTenPercentile(dailyData[3]), 64)
		priceOnHighest, _ := strconv.ParseFloat(replaceTenPercentile(dailyData[4]), 64)
		priceOnLowest, _ := strconv.ParseFloat(replaceTenPercentile(dailyData[5]), 64)
		priceOnClose, _ := strconv.ParseFloat(replaceTenPercentile(dailyData[6]), 64)
		stock.StockNumber = stockNumber
		stock.DealDate = &dealDate
		stock.DealCount = uint64(dealCount)
		stock.PriceOnOpen = priceOnOpen
		stock.PriceOnHighest = priceOnHighest
		stock.PriceOnLowest = priceOnLowest
		stock.PriceOnClose = priceOnClose
		stock.CreatedAt = &now
		stock.UpdatedAt = &now
		r.mongodb.InsertOne(ctx, stock)
	}
	return nil
}

func (r *StockRepository) InsertToMongo(stock *models.StockData) error {
	ctx := context.Background()
	_, err := r.mongodb.InsertOne(ctx, stock)
	if err != nil {
		return err
	}
	return nil
}

func (r *StockRepository) CalcKDVal(stockList []*models.StockData, last *int64) ([]*models.StockKD, error) {
	lastKDList := []*models.StockKD{}
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

		// Asc order
		if int64(i) < *last {
			lastKDList = append(lastKDList, &models.StockKD{
				KVal: lastKVal,
				DVal: lastDVal,
			})
		}
	}
	return lastKDList, nil
}

func (r *StockRepository) Test() {
	ctx := context.Background()
	filter := bson.M{}
	// result, _ := r.mongodb.Find(ctx, filter)
	// fmt.Println(result)
	update := bson.A{
		bson.M{
			"$set": bson.M{
				"deal_date": bson.M{
					"$add": bson.A{
						"$deal_date",
						1000 * 60 * 60 * 8,
					},
				},
			},
		},
	}

	result, err := r.mongodb.UpdateMany(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
