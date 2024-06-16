package handlers

import (
	"exoplanets/helpers"
	"exoplanets/models"
	"exoplanets/sqlOperations"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddExoplanetHandler handles the addition of a new exoplanet
func AddExoplanetHandler(c *gin.Context) {
	var exoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate exoplanet
	if err := helpers.ValidateExoplanet(exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedID, err := sqlOperations.AddExoplanet(exoplanet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf("Exoplanet added successfully with ID %d", insertedID)

	c.JSON(http.StatusCreated, gin.H{"message": message})
}

// ListExoplanetsHandler handles listing all exoplanets
func ListExoplanetsHandler(c *gin.Context) {
	exoplanets, err := sqlOperations.ListExoplanets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exoplanets)
}

// GetExoplanetByIDHandler handles retrieving an exoplanet by its ID
func GetExoplanetByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exoplanet ID"})
		return
	}

	exoplanet, err := sqlOperations.GetExoplanetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exoplanet)
}

// UpdateExoplanetHandler handles updating an exoplanet's details
func UpdateExoplanetHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exoplanet ID"})
		return
	}

	var exoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate exoplanet
	if err := helpers.ValidateExoplanet(exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sqlOperations.UpdateExoplanet(id, exoplanet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exoplanet updated successfully"})
}

// DeleteExoplanetHandler handles deleting an exoplanet
func DeleteExoplanetHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exoplanet ID"})
		return
	}

	if err := sqlOperations.DeleteExoplanet(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exoplanet deleted successfully"})
}

// FuelEstimationHandler handles calculating fuel estimation for a trip to an exoplanet
func FuelEstimationHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exoplanet ID"})
		return
	}

	crewCapacity, err := strconv.Atoi(c.Param("crewCapacity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crew capacity"})
		return
	}

	fuelEstimation, err := sqlOperations.CalculateFuelEstimation(id, crewCapacity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format fuel estimation to two decimal places
	formattedFuelEstimation := fmt.Sprintf("%.2f units", fuelEstimation)

	c.JSON(http.StatusOK, gin.H{"fuel_estimation": formattedFuelEstimation})
}
