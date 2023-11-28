package handlers

import (
	. "TodoListServer/internal/domain/models"
	"golang.org/x/exp/slog"
	"net/http"
)

type IUserService interface {
	GetUser(id uint) (User, error)
}

type RegisterUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func GetUser(log *slog.Logger, userService IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
