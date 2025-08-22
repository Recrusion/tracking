package storage

import "log"

func (td *TrackingDatabase) CreateUser(username, password string) error {
	query := "insert into users (username, password) values ($1, $2)"
	_, err := td.db.Exec(query, username, password)
	if err != nil {
		log.Printf("Ошибка добавления пользователя в базу данных (слой database), %v", err)
	}
	return err
}

func (td *TrackingDatabase) UserVerificationByUsername(username string) (string, error) {
	query := "select id from users where username = $1"
	var id string
	err := td.db.QueryRow(query, username).Scan(&id)
	if err != nil {
		log.Printf("Пользователь с таким username - не найден, продолжаем регистрацию, %v", err)
	}
	return id, nil
}
