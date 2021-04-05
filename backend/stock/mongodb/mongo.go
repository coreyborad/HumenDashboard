package mongodb

import (
	"fmt"
	"log"
	"stock/config"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *mongo.Database

// NewMongoDB NewMongoDB
func NewMongoDB() (err error) {
	var mongoCli *mongo.Client
	if err := config.Check(); err != nil {
		return err
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", config.MongoDB.Host, config.MongoDB.Port))

	// 連接到Mongodb
	mongoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Ping
	err = mongoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	mongoDB = mongoCli.Database(config.MongoDB.DbName)
	return
}

func GetMongoDB() *mongo.Database {
	if mongoDB == nil {
		NewMongoDB()
	}
	return mongoDB
}
