package models

type Indicators struct {
	Username  string `json:"username"`
	Indicator string `json:"indicator"`
	Score     int    `json:"points"`
	Total     int    `json:"max"`
}

type Indicator struct {
	Indicator string `json:"indicator"`
	Score     int    `json:"points"`
	Total     int    `json:"max"`
}
