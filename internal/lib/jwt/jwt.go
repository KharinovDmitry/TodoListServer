package jwt

import (
	"TodoListServer/internal/domain/models"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	ErrInvalidClaims = errors.New("Invalid claims")
)

type UserClaims struct {
	jwt.RegisteredClaims
	ID    uint   `json:"id"`
	Login string `json:"login"`
	Exp   int    `json:"exp"`
}

func NewToken(user models.User, tokenTTL time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["login"] = user.Login
	claims["exp"] = time.Now().Add(tokenTTL).Unix()

	tokenString, err := token.SignedString([]byte("salt")) //TODO
	if err != nil {
		return "", fmt.Errorf("In NewToken: %w", err)
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(_ *jwt.Token) (any, error) { return []byte("salt"), nil },
	)
	if err != nil {
		return nil, fmt.Errorf("In ParseToken: %w", err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("In ParseToken: %w", ErrInvalidClaims)
	}
	return claims, nil
}
