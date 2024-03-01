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

func ListUsers(c *gin.Context) {

	mongoClient := database.GetClient()
	results, err := mongoClient.Database("avito_dev_app").Collection("users").Find(context.TODO(), bson.D{{}})
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

func FindUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	// Find user by id
	var user bson.M

	mongoClient := database.GetClient()
	err = mongoClient.Database("avito_dev_app").Collection("users").FindOne(context.TODO(), bson.D{{"_id", userId}}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// func FindUsers(c *gin.Context) {
// 	userId := c.Param("id")

// 	user := database.FindUser(userId)
// 	if user == nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
// }

// func ListUsers(c *gin.Context) {
// 	listUsers := database.GetUsers()
// 	if listUsers == nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "List is empty"})
// 		return
// 	}

// 	//Response to api
// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "Data fetched successfully", "data": listUsers})
// }
