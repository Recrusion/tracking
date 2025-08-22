package storage

import (
	"context"
	"fmt"
	"log"
)

func (td *TrackingDatabase) GettingPasswordUser(ctx context.Context, username string) (string, error) {
	var password string
	err := td.db.QueryRowContext(ctx, "select password from users where username = $1", username).Scan(&password)
	if err != nil {
		log.Printf("Ошибка получения пароля пользователя по username'у: %v", err)
		return "", fmt.Errorf("error getting password user: %w", err)
	}
	return password, nil
}
