package main

import (
	"os"

	"github.com/Ademayowa/rest-api-demo/db"
	"github.com/Ademayowa/rest-api-demo/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		println("No .env file found, using default environment variables")
	}

	db.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":" + port)
}
