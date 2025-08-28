package models

type Indicator struct {
	Indicator string `json:"indicator"`
	Score     int    `json:"points"`
	Total     int    `json:"max"`
}
