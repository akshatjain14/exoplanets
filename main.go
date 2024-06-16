package main

import (
	"exoplanets/handlers"
	"exoplanets/sqlOperations"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	// Initialize the database connection
	sqlOperations.InitDB()
}

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize the Gin router
	router := gin.Default()

	// Define the routes
	router.POST("/exoplanets", handlers.AddExoplanetHandler)
	router.GET("/exoplanets", handlers.ListExoplanetsHandler)
	router.GET("/exoplanets/:id", handlers.GetExoplanetByIDHandler)
	router.PUT("/exoplanets/:id", handlers.UpdateExoplanetHandler)
	router.DELETE("/exoplanets/:id", handlers.DeleteExoplanetHandler)
	router.GET("/exoplanets/:id/fuel-estimation/:crewCapacity", handlers.FuelEstimationHandler)

	// Start the server
	log.Println("Server starting on port 8080")
	port := "8080"
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
