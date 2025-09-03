package services

import (
	"context"
	"tracking/internal/models"
	"tracking/internal/storage"
)

type Service interface {
	Login(ctx context.Context, username, password string) error
	AddIndicator(ctx context.Context, username, indicator string, total int) error
	IncreaseScore(ctx context.Context, username, indicator string) error
	GetAllIndicators(ctx context.Context, username string) ([]models.Indicator, error)
	DeleteIndicators(ctx context.Context, username, indicator string) error
	CreateAccessToken(ctx context.Context, username string) (string, error)
	CreateRefreshToken(ctx context.Context, username string) (string, error)
	CreateUserService(ctx context.Context, username, password string) error
}
type ServiceTracking struct {
	service storage.Storage
}

func NewServiceTracking(service storage.Storage) *ServiceTracking {
	s := &ServiceTracking{
		service: service,
	}
	return s
}
