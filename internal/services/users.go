package services

import (
	"context"
	"fmt"
)

func (s *ServiceTracking) CreateUserService(ctx context.Context, username, password string) error {
	result, _ := s.service.UserVerificationByUsername(ctx, username)
	if result == "" {
		err := s.service.CreateUser(ctx, username, password)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("a user with this username already exists")
	}
	return nil
}
