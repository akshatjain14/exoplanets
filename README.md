# Exoplanets Management System

The Exoplanets Microservice is a Go-based backend service designed to manage information about exoplanets for space voyagers. It allows users to add, list, retrieve, update, and delete exoplanet data, as well as estimate fuel requirements for trips to specific exoplanets.

## Features

 - Add Exoplanet: Add new exoplanets with properties such as name, description, distance from Earth, radius, 
     mass (for terrestrial planets), and type (GasGiant or Terrestrial).
 - List Exoplanets: Retrieve a list of all available exoplanets.
 - Get Exoplanet by ID: Retrieve detailed information about a specific exoplanet using its unique ID.
 - Update Exoplanet: Update the details of an existing exoplanet.
 - Delete Exoplanet: Remove an exoplanet from the catalog.
 - Fuel Estimation: Calculate the estimated fuel cost for a trip to any exoplanet based on its distance, gravity, and crew capacity.

## Installation

1. Clone the repository:
git clone https://github.com/akshatjain14/exoplanets.git
2. Install dependencies:
go mod tidy
3. Run the application:
go run main.go


The application will start running on port 8080 by default.

## API Endpoints

- `POST /exoplanets` : Add a new exoplanet
- `GET /exoplanets` : Get all exoplanets
- `GET /exoplanets/:id` : Get an exoplanet by ID
- `PUT /exoplanets/:id` : Update an exoplanet by ID
- `DELETE /exoplanets/:id` : Delete an exoplanet by ID
- `GET /exoplanets/:id/fuel-estimation/:crewCapacity` : Calculate fuel estimation for a trip to an exoplanet

## Dependencies

- Gin: Web framework for building APIs in GoLang
- Go-MySQL-Driver: MySQL driver for GoLang






