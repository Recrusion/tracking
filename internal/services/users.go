package services

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *ServiceTracking) CreateUserService(ctx context.Context, username, password string) error {
	result, _ := s.service.UserVerificationByUsername(ctx, username)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if result == "" {
		err = s.service.CreateUser(ctx, username, string(hashedBytes))
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("a user with this username already exists")
	}
	return nil
}
