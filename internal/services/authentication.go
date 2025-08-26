package services

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (s *ServiceTracking) Login(ctx context.Context, username, password string) error {
	passwordDB, err := s.service.GettingPasswordUser(ctx, username)
	if err != nil {
		return err
	}
	if err = checkPasswordHash(password, passwordDB); err != nil {
		return fmt.Errorf("incorrect password")
	}
	return nil
}

func checkPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
