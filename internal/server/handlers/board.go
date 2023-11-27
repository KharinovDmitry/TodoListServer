package handlers

import (
	"TodoListServer/internal/domain/models"
	"TodoListServer/internal/services"
	"TodoListServer/internal/storage/postgres"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

type IBoardService interface {
	CreateBoard(name string, owner uint) error
	DeleteBoard(id uint) error
	GetBoard(id uint, userID uint) (models.Board, error)
	UpdateBoard(id uint, board models.Board) error
}

type CreateBoardRequest struct {
	Name string `json:"name"`
}

func CreateBoard(log *slog.Logger, boardService IBoardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request CreateBoardRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //TODO
			return
		}

		userID, ok := r.Context().Value("claims").(uint)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := boardService.CreateBoard(request.Name, userID)
		if err != nil {
			log.Error(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteBoard(log *slog.Logger, boardService IBoardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func UpdateBoard(log *slog.Logger, boardService IBoardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetBoard(log *slog.Logger, boardService IBoardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Чего-чего?"))
			log.Error(err.Error())
			return
		}

		userID, ok := r.Context().Value("claims").(uint)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		board, err := boardService.GetBoard(uint(id), userID)
		if err != nil {
			if errors.Is(err, postgres.ErrNotFound) {
				w.Write([]byte("Не найдено"))
				//TODO
			}
			if errors.Is(err, services.ErrNotAccess) {
				w.Write([]byte("Нет доступа!"))
				//TODO
			}
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err.Error())
			return
		}

		response, err := json.Marshal(board)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
