package services

import (
	. "TodoListServer/internal/domain/models"
	"TodoListServer/internal/storage"
	"errors"
	"golang.org/x/exp/slog"
)

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

type UserService struct {
	userRepo storage.Repository[User]
	log      *slog.Logger
}

func NewUserService(log *slog.Logger, userRepo storage.Repository[User]) *UserService {
	return &UserService{
		userRepo: userRepo,
		log:      log,
	}
}

func (s *UserService) GetUser(id uint) (User, error) {
	panic("implement me")
}
