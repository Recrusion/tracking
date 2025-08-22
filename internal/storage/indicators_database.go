package storage

import (
	"database/sql"
	"log"
)

func (td *TrackingDatabase) AddIndicator(username, indicator string, total int) error {
	query := "insert into indicators (username, indicator, score, total) values ($1, $2, $3, $4)"
	_, err := td.db.Exec(query, username, indicator, 0, total)
	if err != nil {
		log.Printf("Ошибка добавления цели (слой database), %v", err)
		return err
	}
	return nil
}

func (td *TrackingDatabase) IncreaseScore(username, indicator string) error {
	query := "update indicators set score = score + 1 where username = $1 and indicator = $2"
	_, err := td.db.Exec(query, username, indicator)
	if err != nil {
		log.Printf("Ошибка добавления очков к цели (слой database), %v", err)
		return err
	}
	return nil
}

func (td *TrackingDatabase) GetAllIndicators(username string) (*sql.Rows, error) {
	query := "select indicator, score, total from indicators where username = $1"
	result, err := td.db.Query(query, username)
	if err != nil {
		log.Printf("Ошибка получения всех целей для пользователя %v, (слой database), %v", username, err)
		return nil, err
	}
	return result, nil
}

func (td *TrackingDatabase) DeleteIndicators(username, indicator string) error {
	query := "delete from indicators where username = $1 and indicator = $2"
	_, err := td.db.Exec(query, username, indicator)
	if err != nil {
		log.Printf("Ошибка удаления цели (слой database), %v", err)
		return err
	}
	return nil
}

func (td *TrackingDatabase) GetTotalForIndicator(username, indicator string) (*sql.Rows, error) {
	query := "select total from indicators where username = $1 and indicator = $2"
	result, err := td.db.Query(query, username, indicator)
	if err != nil {
		log.Printf("Ошибка получения конечного числа дней (слой database), %v", err)
		return nil, err
	}
	return result, nil
}

func (td *TrackingDatabase) GetScoreForIndicator(username, indicator string) (*sql.Rows, error) {
	query := "select score from indicators where username = $1 and indicator = $2"
	result, err := td.db.Query(query, username, indicator)
	if err != nil {
		log.Printf("Ошибка получения текущего числа дней (слой database), %v", err)
		return nil, err
	}
	return result, nil
}
