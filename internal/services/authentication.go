package services

import (
	"context"
	"fmt"
)

func (s *ServiceTracking) Login(ctx context.Context, username, password string) error {
	passwordDB, err := s.service.GettingPasswordUser(ctx, username)
	if err != nil {
		return err
	}
	if passwordDB != password {
		return fmt.Errorf("incorrect password")
	}
	return nil
}
