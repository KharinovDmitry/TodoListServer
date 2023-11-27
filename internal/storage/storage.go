package storage

import (
	"TodoListServer/internal/domain/models"
	"TodoListServer/internal/storage/postgres"
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository[T any] interface {
	AddItem(item T) error
	DeleteItem(item T) error
	UpdateItem(id int, newItem T) error
	GetTable() ([]T, error)
	FindItemByID(id int) (T, error)
	FindItemByCondition(condition func(item T) bool) (T, error)
	FindItemsByCondition(condition func(item T) bool) ([]T, error)
}

type Storage struct {
	db      *sql.DB
	Users   Repository[models.User]
	Boards  Repository[models.Board]
	Columns Repository[models.Column]
	Tasks   Repository[models.Task]
}

func New(connStr string) (*Storage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	userRepo := postgres.NewUserRepository(db)
	boardRepo := postgres.NewBoardRepository(db)
	return &Storage{
		db:     db,
		Users:  userRepo,
		Boards: boardRepo,
	}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
