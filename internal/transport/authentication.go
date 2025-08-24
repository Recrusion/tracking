package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"tracking/internal/models"
)

func (h *HandlersTracking) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != "POST" {
		log.Printf("Неправильно выбран метод запроса (должен быть POST)")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	user := models.Users{}
	tokens := make(map[string]string)
	var accessToken, refreshToken string
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Ошибка декодирования данных из тела запроса (слой transport), %v", err)
	}
	err = h.endpoints.Login(ctx, user.Username, user.Password)
	if err == nil {
		accessToken, err = h.endpoints.CreateAccessToken(ctx, user.Username)
		if err != nil {
			log.Printf("Ошибка генерации access-токена (слой transport), %v", err)
		}
		refreshToken, err = h.endpoints.CreateRefreshToken(ctx, user.Username)
		if err != nil {
			log.Printf("Ошибка генерации refresh-токена (слой transport), %v", err)
		}
	} else {
		log.Printf("Неправильный логин или пароль, либо аккаунт не существует (cлой transport), %v", err)
		w.WriteHeader(http.StatusUnauthorized)
	}
	tokens["access_token"] = accessToken
	tokens["refresh_token"] = refreshToken
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tokens)
	if err != nil {
		log.Printf("Ошибка энкодирования токенов, %v", err)
	}
}

func GetUserFromContext(r *http.Request) (string, error) {
	username := r.Context().Value("username")
	if username == nil {
		log.Printf("Username не был передан, либо некорректен")
	}
	return username.(string), nil
}
