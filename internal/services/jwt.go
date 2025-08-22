package services

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func (s *ServiceTracking) CreateAccessToken(ctx context.Context, username string) (string, error) {
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
		return "", err
	}
	err = s.service.SaveAccessToken(ctx, username, tokenString, claims.IssuedAt.Time, claims.ExpiresAt.Time)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *ServiceTracking) CreateRefreshToken(ctx context.Context, username string) (string, error) {
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
		return "", err
	}
	err = s.service.SaveRefreshToken(ctx, username, tokenString, claims.IssuedAt.Time, claims.ExpiresAt.Time)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
