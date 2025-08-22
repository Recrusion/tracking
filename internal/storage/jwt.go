package storage

import (
	"context"
	"fmt"
	"time"
)

func (td *TrackingDatabase) SaveAccessToken(ctx context.Context, username, token string, created, ending time.Time) error {
	_, err := td.db.ExecContext(ctx, "insert into access_tokens (username, access_token, created, ending) values ($1, $2, $3, $4)", username, token, created, ending)
	if err != nil {
		return fmt.Errorf("error saving access token: %w", err)
	}
	return nil
}

func (td *TrackingDatabase) SaveRefreshToken(ctx context.Context, username, token string, created, ending time.Time) error {
	_, err := td.db.ExecContext(ctx, "insert into refresh_tokens (username, refresh_token, created, ending) values ($1, $2, $3, $4)", username, token, created, ending)
	if err != nil {
		return fmt.Errorf("error saving refresh token: %w", err)
	}
	return nil
}
