package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection

func initEngine() {
	var (
		mgoCli   *mongo.Client
		uri      = "you db uri"
		dbName   = "vercel_db"
		collName = "image_random"
		err      error
	)

	clientOptions := options.Client().ApplyURI(uri)

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = mgoCli.Database(dbName).Collection(collName)
}

func GetCollection()  *mongo.Collection{
	if collection == nil {
		initEngine()
	}
	return collection
}

