package services

import (
	"fmt"
	"log"
)

func (s *ServiceTracking) Login(username, password string) error {
	passwordFromDB, err := s.service.GettingPasswordUser(username)
	if err != nil {
		log.Printf("Ошибка получения пароля пользователя или пользователь не существует (слой services), %v", err)
		return err
	}
	if passwordFromDB != password {
		log.Printf("Неправильный пароль (слой services)")
		return fmt.Errorf("incorrect password")
	}
	return nil
}
