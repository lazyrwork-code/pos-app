package main

import (
	"log"

	"pos-app/config"
	"pos-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
