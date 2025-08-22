package database

import (
	"log"
	"time"
)

func (td *TrackingDatabase) SaveAccessToken(username, token string, created, ending time.Time) error {
	query := "insert into access_tokens (username, access_token, created, ending) values ($1, $2, $3, $4)"
	_, err := td.db.Exec(query, username, token, created, ending)
	if err != nil {
		log.Printf("Ошибка сохранения access-токена (слой database), %v", err)
		return err
	}
	return nil
}

func (td *TrackingDatabase) SaveRefreshToken(username, token string, created, ending time.Time) error {
	query := "insert into refresh_tokens (username, refresh_token, created, ending) values ($1, $2, $3, $4)"
	_, err := td.db.Exec(query, username, token, created, ending)
	if err != nil {
		log.Printf("Ошибка сохранения refresh-токена (слой database), %v", err)
		return err
	}
	return nil
}
