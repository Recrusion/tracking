package models

import "time"

type Token struct {
	Username string
	Token    string
	Created  time.Time
	Ending   time.Time
}
