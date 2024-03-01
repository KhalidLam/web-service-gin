package main

import (
	"fmt"
	"log"

	. "github.com/KhalidLam/web-service-gin/database"
	. "github.com/KhalidLam/web-service-gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	fmt.Println("Application is starting....")

	// set env
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Erorr in reading env file", err)
	}
	viper.ReadInConfig()
	apiUrl := viper.GetString("API_URL")

	defer Disconnect()

	// setup gin
	router := gin.Default()
	router.GET("/status", GetStatus)
	router.GET("/users", ListUsers)
	router.GET("/users/:id", FindUser)
	// router.POST("/users", createUser)
	// router.PUT("/users/:id", updateUser)

	router.Run(apiUrl)
}
