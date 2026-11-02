package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mahipalpanwar-7/expense-tracker/internal/config"
)

func main() {
	// Load env file
	err := godotenv.Load()
	if err!=nil{
		log.Println(".env file not found")
	}

	// Connect databse 
	config.ConnectDB()

    // getting port 
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create gin router
	router := gin.Default()

	// Test route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Expense Tracker API is running",
			"status":  "ok",
		})
	})

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "server is healthy",
			"database": "database is connected",
		})
	})

	fmt.Printf("Server running on http://localhost:%s\n", port)

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
