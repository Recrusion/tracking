package services

import (
	"fmt"
	"log"
)

func (s *ServiceTracking) CreateUserService(username, password string) error {
	result, err := s.service.UserVerificationByUsername(username)
	if err != nil {
		log.Printf("Ошибка получения пользователя (слой service), %v", err)
	}
	if result == "" {
		err = s.service.CreateUser(username, password)
		if err != nil {
			log.Printf("Ошибка создания пользователя (слой service), %v", err)
		}
	} else {
		return fmt.Errorf("пользователь с таким username уже существует")
	}
	return nil
}
