package services

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func (s *ServiceTracking) CreateAccessToken(username string) (string, error) {
	accessSecret := os.Getenv("TOKEN_SECRET")
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		log.Printf("Ошибка создания access-токена (слой services), %v", err)
		return "", err
	}
	err = s.service.SaveAccessToken(username, tokenString, claims.IssuedAt.Time, claims.ExpiresAt.Time)
	if err != nil {
		log.Printf("Ошибка сохранения access-токена (слой services), %v", err)
		return "", err
	}
	return tokenString, nil
}

func (s *ServiceTracking) CreateRefreshToken(username string) (string, error) {
	refreshSecret := os.Getenv("TOKEN_SECRET")
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * (24 * time.Hour))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(refreshSecret))
	if err != nil {
		log.Printf("Ошибка создания refresh-токена (слой services), %v", err)
		return "", err
	}
	err = s.service.SaveRefreshToken(username, tokenString, claims.IssuedAt.Time, claims.ExpiresAt.Time)
	if err != nil {
		log.Printf("Ошибка сохранения refresh-токена (слой services), %v", err)
		return "", err
	}
	return tokenString, nil
}
