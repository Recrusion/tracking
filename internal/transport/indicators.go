package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Indicators struct {
	Indicator string `json:"indicator"`
	Total     int    `json:"total"`
}

func (h *HandlersTracking) AddIndicator(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != "PUT" {
		log.Printf("Неправильно выбран метод запроса (должен быть PUT)")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	username, err := GetUserFromContext(r)
	if err != nil {
		log.Printf("Ошибка получения username'a пользователя из контекста реквеста (слой transport), %v", err)
	}
	indicator := Indicators{}
	err = json.NewDecoder(r.Body).Decode(&indicator)
	if err != nil {
		log.Printf("Ошибка декодирования данных из тела запроса (слой transport), %v", err)
	}
	err = h.endpoints.AddIndicator(ctx, username, indicator.Indicator, indicator.Total)
	if err != nil {
		log.Printf("Ошибка добавления цели (слой transport), %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *HandlersTracking) IncreaseScore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != "PUT" {
		log.Printf("Неправильно выбран метод запроса (должен быть PUT)")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	username, err := GetUserFromContext(r)
	if err != nil {
		log.Printf("Ошибка получения username'a пользователя из контекста реквеста (слой transport), %v", err)
	}
	indicator := Indicators{}
	err = json.NewDecoder(r.Body).Decode(&indicator)
	if err != nil {
		log.Printf("Ошибка декодирования данных из тела запроса (слой transport), %v", err)
	}
	err = h.endpoints.IncreaseScore(ctx, username, indicator.Indicator)
	if err != nil {
		log.Printf("Ошибка добавления очков к цели (слой transport), %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HandlersTracking) GetAllIndicators(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != "GET" {
		log.Printf("Неправильно выбран метод запроса (должен быть GET)")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	username, err := GetUserFromContext(r)
	if err != nil {
		log.Printf("Ошибка получения username'a пользователя из контекста реквеста (слой transport), %v", err)
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		log.Printf("Ошибка получения количества строк из URL-адреса (слой transport), %v", err)
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		log.Printf("Ошибка получения номера страницы из URL-адреса (слой transport), %v", err)
	}

	if pageSize < 1 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}

	indicators, err := h.endpoints.GetAllIndicators(ctx, username, pageSize, page)
	if err != nil {
		log.Printf("Ошибка получения всех целей для пользователя %v, (слой transport), %v", username, err)
	}
	err = json.NewEncoder(w).Encode(indicators)
	if err != nil {
		log.Printf("Ошибка энкодирования целей, %v", err)
	}
	response := map[string]interface{}{
		"pageSize": pageSize,
		"page":     page,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Ошибка энкодирования данных о странице, %v", err)
	}
}

func (h *HandlersTracking) DeleteIndicators(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != "DELETE" {
		log.Printf("Неправильно выбран метод запроса (должен быть DELETE)")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json")
	username, err := GetUserFromContext(r)
	if err != nil {
		log.Printf("Ошибка получения username'a пользователя из контекста реквеста (слой transport), %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	indicator := Indicators{}
	err = json.NewDecoder(r.Body).Decode(&indicator)
	if err != nil {
		log.Printf("Ошибка декодирования данных из тела запроса (слой transport), %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.endpoints.DeleteIndicators(ctx, username, indicator.Indicator)
	if err != nil {
		log.Printf("Ошибка удаления цели (слой transport), %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
