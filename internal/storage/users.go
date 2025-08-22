package storage

import (
	"context"
	"fmt"
)

func (td *TrackingDatabase) CreateUser(ctx context.Context, username, password string) error {
	_, err := td.db.ExecContext(ctx, "insert into users (username, password) values ($1, $2)", username, password)
	if err != nil {
		return fmt.Errorf("error create user: %w", err)
	}
	return nil
}

func (td *TrackingDatabase) UserVerificationByUsername(ctx context.Context, username string) (string, error) {
	var id string
	err := td.db.SelectContext(ctx, &id, "select id from users where username = $1", username)
	if err != nil {
		return "", fmt.Errorf("error get user or user does not exist: %w", err)
	}
	return id, nil
}
