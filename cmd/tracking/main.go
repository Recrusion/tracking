package main

import (
	"log"
	"net/http"
	"tracking/internal/database"
	"tracking/internal/middleware"
	"tracking/internal/services"
	"tracking/internal/transport"

	"github.com/pressly/goose"
	"github.com/rs/cors"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Printf("Ошибка подключения к базе данных, %v", err)
	} else {
		log.Printf("Подключение к базе данных - успешно!")
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Ошибка проверки подключения к базе данных, %v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Printf("Ошибка выбора диалекта, %v", err)
	}

	if err := goose.Up(db, "/internal/migrations"); err != nil {
		log.Printf("Ошибка запуска миграций базы данных, %v", err)
	}

	database := database.NewTrackingDatabase(db)
	service := services.NewServiceTracking(database)
	handler := transport.NewHandlersTracking(service)

	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/public/index.html", http.StatusMovedPermanently)
	})
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/register.html")
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/login.html")
	})
	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/dashboard.html")
	})

	mux.HandleFunc("/api/register", handler.CreateUser)
	mux.HandleFunc("/api/login", handler.Login)
	mux.HandleFunc("/api/addindicators", middleware.AuthenticationMiddleware(handler.AddIndicator))
	mux.HandleFunc("/api/increase", middleware.AuthenticationMiddleware(handler.IncreaseScore))
	mux.HandleFunc("/api/getallindicators", middleware.AuthenticationMiddleware(handler.GetAllIndicators))
	mux.HandleFunc("/api/deleteindicator", middleware.AuthenticationMiddleware(handler.DeleteIndicators))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	cors := c.Handler(mux)

	log.Printf("Сервер прослушивается на порту :8080...")
	err = http.ListenAndServe(":8080", cors)
	if err != nil {
		log.Printf("Ошибка запуска сервера, %v", err)
	}
}
