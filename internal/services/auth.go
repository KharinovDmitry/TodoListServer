package services

import (
	. "TodoListServer/internal/domain/models"
	"TodoListServer/internal/lib/jwt"
	"TodoListServer/internal/storage"
	"TodoListServer/internal/storage/postgres"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
	"time"
)

type AuthService struct {
	userRepo storage.Repository[User]
	log      *slog.Logger
	tokenTTL time.Duration
}

func NewAuthService(log *slog.Logger, userRepo storage.Repository[User], tokenTTL time.Duration) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		log:      log,
		tokenTTL: tokenTTL,
	}
}

func (s *AuthService) RegisterUser(login string, password string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("In UserService(RegisterUser): %w", err)
	}

	err = s.userRepo.AddItem(User{Login: login, Password: passHash})
	if err != nil {
		return fmt.Errorf("In UserService(RegisterUser): %w", err)
	}

	return nil
}

func (s *AuthService) LoginUser(login string, password string) (string, error) {
	user, err := s.userRepo.FindItemByCondition(
		func(item User) bool {
			return item.Login == login
		})

	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return "", fmt.Errorf("In UserService(LoginUser): %w", err)
		}

		return "", fmt.Errorf("In UserService(LoginUser): %w", err)
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return "", fmt.Errorf("In UserService(LoginUser): %w", ErrInvalidCredentials)
	}

	token, err := jwt.NewToken(user, s.tokenTTL)
	if err != nil {
		return "", fmt.Errorf("In UserService(LoginUser): %w", err)
	}

	return token, nil
}
