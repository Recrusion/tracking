package storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"tracking/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage interface {
	GettingPasswordUser(ctx context.Context, username string) (string, error)
	AddIndicator(ctx context.Context, username, indicator string, total int) error
	DeleteIndicators(ctx context.Context, username, indicator string) error
	IncreaseScore(ctx context.Context, username, indicator string) error
	GetAllIndicators(ctx context.Context, username string) ([]models.Indicator, error)
	GetTotalForIndicator(ctx context.Context, username, indicator string) (int, error)
	GetScoreForIndicator(ctx context.Context, username, indicator string) (int, error)
	SaveAccessToken(ctx context.Context, username, token string, created, ending time.Time) error
	SaveRefreshToken(ctx context.Context, username, token string, created, ending time.Time) error
	CreateUser(ctx context.Context, username, password string) error
	UserVerificationByUsername(ctx context.Context, username string) (string, error)
}

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
