package helpers

import (
	"errors"
	"exoplanets/models"
	"math"
)

// ValidateExoplanet validates the properties of an exoplanet
func ValidateExoplanet(exoplanet models.Exoplanet) error {
	if exoplanet.Distance <= 10 || exoplanet.Distance >= 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius <= 0.1 || exoplanet.Radius >= 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if exoplanet.TypeName == "Terrestrial" && (*exoplanet.Mass <= 0.1 || *exoplanet.Mass >= 10) {
		return errors.New("mass for terrestrial exoplanet must be between 0.1 and 10 Earth-mass units")
	}
	return nil
}

// CalculateGravity calculates the gravity of an exoplanet
func CalculateGravity(exoplanet models.Exoplanet) float64 {
	if exoplanet.TypeName == "GasGiant" {
		return 0.5 / math.Pow(exoplanet.Radius, 2)
	}
	if exoplanet.TypeName == "Terrestrial" {
		return *exoplanet.Mass / math.Pow(exoplanet.Radius, 2)
	}
	return 0.0
}
