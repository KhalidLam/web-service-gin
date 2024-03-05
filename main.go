package main

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/KhalidLam/web-service-gin/database"
	. "github.com/KhalidLam/web-service-gin/database"
	. "github.com/KhalidLam/web-service-gin/routes"
	"github.com/KhalidLam/web-service-gin/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	fmt.Println("Application is starting....")

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	apiUrl := config.ServerAddress

	defer Disconnect()

	// GenerateFakeData()

	// setup gin
	router := gin.Default()
	router.GET("/status", GetStatus)
	router.GET("/transactions", ListTransactions)
	router.GET("/transactions/:id", FindTransaction)
	router.POST("/transactions", CreateTransaction)
	router.PUT("/transactions/:id", CancelTransaction)

	router.Run(apiUrl)
}

func GenerateFakeData() {

	mongoClient := database.GetClient()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	quickstartDatabase := mongoClient.Database("avito_dev_app")
	usersCollection := quickstartDatabase.Collection("users")
	productsCollection := quickstartDatabase.Collection("products")
	transactionCollection := quickstartDatabase.Collection("transactions")

	userResult, err := usersCollection.InsertOne(ctx, bson.D{
		{"name", "Jhon Smith"},
		{"email", "test@email.com"},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted document into user collection:  %v \n", userResult.InsertedID)

	productsResult, err := productsCollection.InsertOne(ctx, bson.D{
		{"name", "Product 1"},
		{"price", 78},
		{"description", "This is a description..."},
		{"stocks", 5},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted document into product collection:  %v \n", productsResult.InsertedID)

	transactionsResult, err := transactionCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"productId", productsResult.InsertedID},
			{"userId", userResult.InsertedID},
			{"total", 250},
			{"quantity", 1},
			{"date", time.Now().Unix()},
			{"status", "pending"},
		},
		bson.D{
			{"productId", productsResult.InsertedID},
			{"userId", userResult.InsertedID},
			{"total", 600},
			{"quantity", 2},
			{"date", time.Now().Unix()},
			{"status", "pending"},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted %v documents into transaction collection!\n", len(transactionsResult.InsertedIDs))
}
