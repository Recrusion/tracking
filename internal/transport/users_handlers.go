package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"tracking/internal/models"
)

func (h *HandlersTracking) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Неправильный метод запроса (должен быть POST)")
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Printf("Заголовок Content-Type должен быть application/json")
	}

	user := models.Users{
		Username: "",
		Password: "",
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Ошибка декодирования данных из тела запроса, %v", err)
	}

	err = h.handlers.CreateUserService(user.Username, user.Password)
	if err != nil {
		log.Printf("Ошибка создания пользователя (слой transport), %v", err)
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
