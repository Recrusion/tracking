package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"tracking/internal/middleware"
	"tracking/internal/services"
	"tracking/internal/storage"
	"tracking/internal/transport"

	"github.com/rs/cors"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Ошибка загрузки данных из переменной окружения, %v", err)
		log.Fatalln(err)
	}

	db, err := storage.InitDB(ctx)
	if err != nil {
		log.Printf("Ошибка подключения к базе данных, %v", err)
		log.Fatalln(err)
	}

	defer db.Close()

	database := storage.NewTrackingDatabase(db)
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
