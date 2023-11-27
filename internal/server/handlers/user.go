package handlers

import (
	. "TodoListServer/internal/domain/models"
	"encoding/json"
	"golang.org/x/exp/slog"
	"net/http"
)

type IUserService interface {
	RegisterUser(login string, password string) error
	LoginUser(login string, password string) (string, error)
	GetUser(id uint) (User, error)
}

type RegisterUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func RegisterUser(log *slog.Logger, userService IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RegisterUserRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //TODO
			return
		}
		if err := userService.RegisterUser(request.Login, request.Password); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //TODO
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func LoginUser(log *slog.Logger, userService IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RegisterUserRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //TODO
			return
		}
		token, err := userService.LoginUser(request.Login, request.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //TODO
			return
		}
		w.Write([]byte(token))
		w.WriteHeader(http.StatusOK)
	}
}

func GetUser(log *slog.Logger, userService IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
