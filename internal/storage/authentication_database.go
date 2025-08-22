package database

import "log"

func (td *TrackingDatabase) GettingPasswordUser(username string) (string, error) {
	var password string
	query := "select password from users where username = $1"
	err := td.db.QueryRow(query, username).Scan(&password)
	if err != nil {
		log.Printf("Ошибка получения пароля пользователя или пользователь не существует (слой database), %v", err)
		return "", err
	}
	return password, nil
}
