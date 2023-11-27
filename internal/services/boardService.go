package services

import (
	. "TodoListServer/internal/domain/models"
	"TodoListServer/internal/storage"
	"errors"
	"fmt"
	"golang.org/x/exp/slog"
)

var (
	ErrNotAccess = errors.New("Not access")
)

type BoardService struct {
	boardRepo storage.Repository[Board]
	log       *slog.Logger
}

func NewBoardService(log *slog.Logger, repo storage.Repository[Board]) *BoardService {
	return &BoardService{
		boardRepo: repo,
		log:       log,
	}
}

func (s *BoardService) CreateBoard(name string, owner uint) error {
	err := s.boardRepo.AddItem(Board{
		OwnerID: owner,
		Name:    name,
	})
	if err != nil {
		err = fmt.Errorf("In BoardService(CreateBoard): %w", err)
	}
	return err
}

func (s *BoardService) DeleteBoard(id uint) error {
	panic("not implemented")
}

func (s *BoardService) UpdateBoard(id uint, board Board) error {
	panic("not implemented")
}

func (s *BoardService) GetBoard(id uint, userID uint) (Board, error) {
	board, err := s.boardRepo.FindItemByID(int(id))
	if err != nil {
		return Board{}, fmt.Errorf("In BoardService(GetBoard): %w", err)
	}

	if board.OwnerID != userID { //TODO
		return Board{}, fmt.Errorf("In BoardService(GetBoard): %w", ErrNotAccess)
	}

	return board, err
}
