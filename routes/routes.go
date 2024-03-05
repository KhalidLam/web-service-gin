package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/KhalidLam/web-service-gin/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Working")
}

func ListTransactions(c *gin.Context) {

	mongoClient := database.GetClient()
	results, err := mongoClient.Database("avito_dev_app").Collection("transactions").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var users []bson.M
	if err = results.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func FindTransaction(c *gin.Context) {
	id := c.Param("id")
	transactionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	var transaction bson.M

	mongoClient := database.GetClient()
	err = mongoClient.Database("avito_dev_app").Collection("transactions").FindOne(context.TODO(), bson.D{{"_id", transactionId}}).Decode(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func CreateTransaction(c *gin.Context) {

	var newTransaction bson.M

	if err := c.BindJSON(&newTransaction); err != nil {
		return
	}

	mongoClient := database.GetClient()
	transactionResult, err := mongoClient.Database("avito_dev_app").Collection("transactions").InsertOne(c, newTransaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactionResult)
}

func CancelTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction bson.M

	transactionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	mongoClient := database.GetClient()
	if err := mongoClient.Database("avito_dev_app").Collection("transactions").FindOne(c, bson.M{"_id": transactionId}).Decode(&transaction); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// change transaction status to canceled
	updateResult, err := mongoClient.Database("avito_dev_app").Collection("transactions").UpdateOne(c, bson.D{{"_id", transactionId}}, bson.D{{"$set", bson.D{{"status", "canceled"}}}})

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, updateResult)

}
