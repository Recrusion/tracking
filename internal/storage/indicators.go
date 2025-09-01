package storage

import (
	"context"
	"fmt"
	"tracking/internal/models"
)

func (td *TrackingDatabase) AddIndicator(ctx context.Context, username, indicator string, total int) error {
	_, err := td.db.ExecContext(ctx, "insert into indicators (username, indicator, score, total) values ($1, $2, $3, $4)", username, indicator, 0, total)
	if err != nil {
		return fmt.Errorf("error add indicator: %w", err)
	}
	return nil
}

func (td *TrackingDatabase) IncreaseScore(ctx context.Context, username, indicator string) error {
	_, err := td.db.ExecContext(ctx, "update indicators set score = score + 1 where username = $1 and indicator = $2", username, indicator)
	if err != nil {
		return fmt.Errorf("error increase score: %w", err)
	}
	return nil
}

func (td *TrackingDatabase) GetAllIndicators(ctx context.Context, username string, limit, offset int) ([]models.Indicator, error) {
	var indicators []models.Indicator
	err := td.db.SelectContext(ctx, &indicators, "select indicator, score, total from indicators where username = $1 order by indicator desc limit $1 offset $2", username, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error get all indicators: %w", err)
	}
	return indicators, nil
}

func (td *TrackingDatabase) DeleteIndicators(ctx context.Context, username, indicator string) error {
	_, err := td.db.ExecContext(ctx, "delete from indicators where username = $1 and indicator = $2", username, indicator)
	if err != nil {
		return fmt.Errorf("error delete indicator: %w", err)
	}
	return nil
}

func (td *TrackingDatabase) GetTotalForIndicator(ctx context.Context, username, indicator string) (int, error) {
	var total int
	err := td.db.GetContext(ctx, &total, "select total from indicators where username = $1 and indicator = $2", username, indicator)
	if err != nil {
		return 0, fmt.Errorf("error get total for indicator is %s: %w", indicator, err)
	}
	return total, nil
}

func (td *TrackingDatabase) GetScoreForIndicator(ctx context.Context, username, indicator string) (int, error) {
	var score int
	err := td.db.GetContext(ctx, &score, "select score from indicators where username = $1 and indicator = $2", username, indicator)
	if err != nil {
		return 0, fmt.Errorf("error get score for indicator is %s: %w", indicator, err)
	}
	return score, nil
}
