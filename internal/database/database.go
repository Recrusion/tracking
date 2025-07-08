package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type TrackingDatabase struct {
	db *sql.DB
}

func NewTrackingDatabase(db *sql.DB) *TrackingDatabase {
	td := &TrackingDatabase{
		db: db,
	}
	return td
}

func InitDB() (*sql.DB, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Ошибка загрузки данных из переменной окружения, %v", err)
	}

	connectString := os.Getenv("CONNECT_DATABASE")
	database, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Printf("Ошибка подключения к базе данных, %v", err)
	}

	err = database.Ping()
	if err != nil {
		log.Printf("Ошибка проверки подключения к базе данных, %v", err)
	}

	return database, nil
}
