package services

import (
	"context"
	"fmt"
	"tracking/internal/models"
)

func (s *ServiceTracking) AddIndicator(ctx context.Context, username, indicator string, total int) error {
	err := s.service.AddIndicator(ctx, username, indicator, total)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceTracking) IncreaseScore(ctx context.Context, username, indicator string) error {
	err := s.service.IncreaseScore(ctx, username, indicator)
	if err != nil {
		return err
	}
	total, err := s.service.GetTotalForIndicator(ctx, username, indicator)
	if err != nil {
		return err
	}

	score, err := s.service.GetScoreForIndicator(ctx, username, indicator)
	if err != nil {
		return err
	}

	if score > total {
		return fmt.Errorf("score cannot be greater than total")
	}
	return nil
}

func (s *ServiceTracking) GetAllIndicators(ctx context.Context, username string, pageSize, page int) ([]models.Indicator, error) {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	result, err := s.service.GetAllIndicators(ctx, username, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceTracking) DeleteIndicators(ctx context.Context, username, indicator string) error {
	err := s.service.DeleteIndicators(ctx, username, indicator)
	if err != nil {
		return err
	}
	return nil
}
