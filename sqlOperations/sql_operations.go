package sqlOperations

import (
	"database/sql"
	"exoplanets/models"
	"fmt"
	"log"
	"math"

	"exoplanets/helpers"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	// Initialize database connection
	dsn := "root:abc@123@tcp(localhost:3306)/exoplanets"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Connected to MySQL database")
}

func AddExoplanet(exoplanet models.Exoplanet) (int, error) {
	query := `INSERT INTO exoplanets (name, description, distance, radius, mass, type_name) 
				VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, exoplanet.Name, exoplanet.Description, exoplanet.Distance, exoplanet.Radius, exoplanet.Mass, exoplanet.TypeName)
	if err != nil {
		return 0, fmt.Errorf("error inserting exoplanet: %v", err)
	}

	// Get the ID of the inserted exoplanet
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %v", err)
	}

	return int(id), nil
}

func ListExoplanets() ([]models.Exoplanet, error) {
	query := `SELECT * FROM exoplanets`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching exoplanets: %v", err)
	}
	defer rows.Close()

	var exoplanets []models.Exoplanet
	for rows.Next() {
		var exoplanet models.Exoplanet
		err := rows.Scan(&exoplanet.ID, &exoplanet.Name, &exoplanet.Description, &exoplanet.Distance, &exoplanet.Radius, &exoplanet.Mass, &exoplanet.TypeName)
		if err != nil {
			return nil, fmt.Errorf("error scanning exoplanet row: %v", err)
		}
		exoplanets = append(exoplanets, exoplanet)
	}

	return exoplanets, nil
}

func GetExoplanetByID(id int) (models.Exoplanet, error) {
	var exoplanet models.Exoplanet
	query := `SELECT * FROM exoplanets WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&exoplanet.ID, &exoplanet.Name, &exoplanet.Description, &exoplanet.Distance, &exoplanet.Radius, &exoplanet.Mass, &exoplanet.TypeName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Exoplanet{}, fmt.Errorf("exoplanet with ID %d not found", id)
		}
		return models.Exoplanet{}, fmt.Errorf("error fetching exoplanet: %v", err)
	}
	return exoplanet, nil
}

func UpdateExoplanet(id int, exoplanet models.Exoplanet) error {
	query := `UPDATE exoplanets SET name = ?, description = ?, distance = ?, radius = ?, mass = ?, type_name = ? WHERE id = ?`
	_, err := db.Exec(query, exoplanet.Name, exoplanet.Description, exoplanet.Distance, exoplanet.Radius, exoplanet.Mass, exoplanet.TypeName, id)
	if err != nil {
		return fmt.Errorf("error updating exoplanet: %v", err)
	}
	return nil
}

func DeleteExoplanet(id int) (int64, error) {
	query := `DELETE FROM exoplanets WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("error deleting exoplanet: %v", err)
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error getting rows affected: %v", err)
	}

	return rowsAffected, nil
}

func CalculateFuelEstimation(id int, crewCapacity int) (float64, error) {
	exoplanet, err := GetExoplanetByID(id)
	if err != nil {
		return 0.0, fmt.Errorf("error getting exoplanet details: %v", err)
	}

	// Calculate gravity for the exoplanet
	gravity := helpers.CalculateGravity(exoplanet)

	// Calculate fuel estimation
	fuelEstimation := exoplanet.Distance / (math.Pow(gravity, 2)) * float64(crewCapacity)

	return fuelEstimation, nil
}
