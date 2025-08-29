package models

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Indicator struct {
	Indicator string `json:"indicator"`
	Score     int    `json:"points"`
	Total     int    `json:"max"`
}
