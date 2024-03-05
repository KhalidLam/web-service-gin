package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/KhalidLam/web-service-gin/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetClient() *mongo.Client {
	if client != nil {
		return client
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	uri := config.DBSource

	if uri == "" {
		log.Fatal("mongodb uri string not found!")
	}

	// getting client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("MongoClient connected..")

	return client
}

func GetCollection(client *mongo.Client, collectioName string) *mongo.Collection {
	collection := client.Database("avito_dev_app").Collection(collectioName)
	return collection
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if client == nil {
		return
	}
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
