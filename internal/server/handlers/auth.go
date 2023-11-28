package handlers

import (
	"encoding/json"
	"golang.org/x/exp/slog"
	"net/http"
)

type IAuthService interface {
	RegisterUser(login string, password string) error
	LoginUser(login string, password string) (string, error)
}

func RegisterUser(log *slog.Logger, authService IAuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RegisterUserRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //TODO
			return
		}
		if err := authService.RegisterUser(request.Login, request.Password); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //TODO
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func LoginUser(log *slog.Logger, authService IAuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RegisterUserRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //TODO
			return
		}
		token, err := authService.LoginUser(request.Login, request.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //TODO
			return
		}
		w.Write([]byte(token))
		w.WriteHeader(http.StatusOK)
	}
}
