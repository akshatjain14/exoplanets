package models

type Exoplanet struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Distance    float64  `json:"distance"`
	Radius      float64  `json:"radius"`
	Mass        *float64 `json:"mass,omitempty"`
	TypeName    string   `json:"type_name"`
}
