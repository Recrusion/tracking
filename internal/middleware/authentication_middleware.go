package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"tracking/internal/services"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

var usernameKey contextKey = "username"

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tokenString string
		header := r.Header.Get("Authorization")
		details := strings.Split(header, " ")
		if len(details) != 2 || strings.ToUpper(details[0]) != "BEARER" {
			log.Printf("Неправильный формат токена, либо токен поврежден или не передан")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString = details[1]
		token, err := jwt.ParseWithClaims(tokenString, &services.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := fmt.Errorf("mddlwr: Unexpected token signature")
				w.WriteHeader(http.StatusUnauthorized)
				return nil, err
			} else {
				return []byte(os.Getenv("TOKEN_SECRET")), nil
			}
		})
		if err != nil {
			log.Printf("Неизвестная подпись токена (слой middleware), %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*services.Claims)
		if ok && token.Valid {
			if time.Now().After(claims.ExpiresAt.Time) {
				err = fmt.Errorf("mddwlr: token is not valid")
				log.Printf("Срок действия токена истек (слой middleware), %v", err)
				w.WriteHeader(http.StatusUnauthorized)
			}
		}
		ctx := context.WithValue(r.Context(), usernameKey, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
