package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"tracking/internal/models"
)

type MockHandlers struct {
	CreateUserServiceFunc func(username, password string) error
}

type TestHandlers struct {
	handlers *MockHandlers
}

func (t *TestHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type not supported", http.StatusUnsupportedMediaType)
		return
	}

	user := models.Users{
		Username: "",
		Password: "",
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = t.handlers.CreateUserServiceFunc(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func TestCreateUser(t *testing.T) {
	mockHandlers := &MockHandlers{
		CreateUserServiceFunc: func(username, password string) error {
			if username != "testuser" {
				return nil
			}
			return fmt.Errorf("пользователь уже существует или данные не верны")
		},
	}

	handlers := &TestHandlers{handlers: mockHandlers}

	// случай когда пользователь существует
	user := models.Users{Username: "testuser", Password: "testpass"}
	jsonData, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateUser(w, req)

	if status := w.Code; status != http.StatusConflict {
		t.Errorf("Ожидали 409, а пришёл %d", status)
	}

	// случай когда пользователя не существует

	user = models.Users{Username: "testuser1", Password: "testpass"}
	jsonData, err = json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	handlers.CreateUser(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Ожидали 201, а пришёл %d", status)
	}

}
