package storage

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

var (
	dbOnce     sync.Once
	dbInstance *sqlx.DB
	onceError  error
)

type TrackingDatabase struct {
	db *sqlx.DB
}

func NewTrackingDatabase(db *sqlx.DB) *TrackingDatabase {
	td := &TrackingDatabase{
		db: db,
	}
	return td
}

func InitDB(ctx context.Context) (*sqlx.DB, error) {
	dbOnce.Do(func() {
		db, err := sqlx.ConnectContext(ctx, "postgres", os.Getenv("DSN"))
		if err != nil {
			log.Printf("Ошибка подключения к базе данных tracking: %v", err)
			onceError = fmt.Errorf("error connecting to database: %w", err)
		}
		dbInstance = db
		onceError = nil
	})
	return dbInstance, onceError
}
