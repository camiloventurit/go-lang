package main

import (
	"log"
	"myApi/config"
	"myApi/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, error := config.DBConnection()
	sqlDB, err := db.DB()

	if error != nil {
		log.Fatalf("Failed to connect to database : %v", err)
	}

	// Set up API routes
	api := r.Group("/api")
	routes.SetupRoutes(api, db)

	// Define a custom 404 handler
	r.NoRoute(func(ctx *gin.Context) {
		urlRequested := ctx.Request.URL
		ctx.JSON(http.StatusNotFound, gin.H{
			"message":       "Welcome to the api :) this routes doesn't exists",
			"url_requested": urlRequested.Path,
		})
	})

	defer sqlDB.Close()

	// Start server
	r.Run(":8080")
}
