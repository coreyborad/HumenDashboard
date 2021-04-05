package repositories

import (
	"context"
	"errors"
	"stock/models"
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
func (r *StockRepository) Find(id uint64) (*models.Stock, error) {
	stock := new(models.Stock)
	if err := r.db.First(&stock, id).Error; err != nil {
		return nil, err
	}

	return stock, nil
}

// Find Find
func (r *StockRepository) GetStockDataByDate(stock_number string, start_date *time.Time, end_date *time.Time) ([]*models.StockData, error) {
	ctx := context.Background()
	filter := bson.M{
		"stock_number": stock_number,
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
