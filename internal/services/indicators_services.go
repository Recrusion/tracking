package services

import (
	"fmt"
	"log"
	"tracking/internal/models"
)

func (s *ServiceTracking) AddIndicator(username, indicator string, total int) error {
	err := s.service.AddIndicator(username, indicator, total)
	if err != nil {
		log.Printf("Ошибка добавления цели (слой services), %v", err)
		return err
	}
	return nil
}

func (s *ServiceTracking) IncreaseScore(username, indicator string) error {
	err := s.service.IncreaseScore(username, indicator)
	if err != nil {
		log.Printf("Ошибка добавления очков к цели (слой services), %v", err)
		return err
	}
	result, err := s.service.GetTotalForIndicator(username, indicator)
	if err != nil {
		log.Printf("Ошибка получения конечного числа дней (слой services), %v", err)
	}
	var total int
	for result.Next() {
		err = result.Scan(&total)
		if err != nil {
			log.Printf("Ошибка сканирования конечного числа дней из ответа базы данных (слой services), %v", err)
		}
	}
	result, err = s.service.GetScoreForIndicator(username, indicator)
	if err != nil {
		log.Printf("Ошибка получения текущего числа дней (слой services), %v", err)
	}
	var score int
	for result.Next() {
		err = result.Scan(&score)
		if err != nil {
			log.Printf("Ошибка сканирования конечного числа дней из ответа базы данных (слой services), %v", err)
		}
	}

	if score > total {
		err = fmt.Errorf("score не может быть больше total")
		log.Printf("Цель достигнута, стрик достиг своего максимума, %v", err)
		return err
	}
	return nil
}

func (s *ServiceTracking) GetAllIndicators(username string) ([]models.Indicator, error) {
	result, err := s.service.GetAllIndicators(username)
	if err != nil {
		log.Printf("Ошибка получения всех целей для пользователя %v, (слой services), %v", username, err)
		return nil, err
	}
	indicators := []models.Indicator{}
	for result.Next() {
		i := models.Indicator{}
		err = result.Scan(&i.Indicator, &i.Score, &i.Total)
		if err != nil {
			log.Printf("Ошибка сканирования данных из ответа базы данных (слой services), %v", err)
		}
		indicators = append(indicators, i)
	}
	return indicators, nil
}

func (s *ServiceTracking) DeleteIndicators(username, indicator string) error {
	err := s.service.DeleteIndicators(username, indicator)
	if err != nil {
		log.Printf("Ошибка удаления цели (слой services), %v", err)
		return err
	}
	return nil
}
