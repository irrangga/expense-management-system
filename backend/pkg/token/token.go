package token

import (
	"backend/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token interface {
	GenerateJWT(userID int64, email, role string) (string, error)
}

type token struct{}

func NewToken() Token {
	return &token{}
}

type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (t *token) GenerateJWT(userID int64, email, role string) (string, error) {
	now := time.Now()
	expirationTime := now.Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.Auth.SecretKey))
}
