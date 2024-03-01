package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
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

	uri := viper.GetString("MONGODB_URI")
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
