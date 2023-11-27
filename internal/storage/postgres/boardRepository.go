package postgres

import (
	. "TodoListServer/internal/domain/models"
	"database/sql"
	"errors"
	"fmt"
)

type BoardRepository struct {
	db *sql.DB
}

func NewBoardRepository(db *sql.DB) *BoardRepository {
	return &BoardRepository{db: db}
}

func (r *BoardRepository) AddItem(item Board) error {
	_, err := r.db.Exec("INSERT INTO boards(OwnerId, Name) VALUES ($1, $2)",
		item.OwnerID, item.Name)
	if err != nil {
		err = fmt.Errorf("In BoardRepository(AddItem): %w", err)
	}
	return err
}

func (r *BoardRepository) DeleteItem(item Board) error {
	panic("not implemented")
}

func (r *BoardRepository) UpdateItem(id int, newItem Board) error {
	panic("not implemented")
}

func (r *BoardRepository) GetTable() ([]Board, error) {
	rows, err := r.db.Query("SELECT * FROM boards")
	if err != nil {
		return nil, fmt.Errorf("In BoardRepository(GetTable): %w", err)
	}
	defer rows.Close()

	users := make([]Board, 0) //??
	for rows.Next() {
		var board Board
		err = rows.Scan(&board.ID, &board.OwnerID, &board.Name)
		if err != nil {
			return nil, fmt.Errorf("In BoardRepository(GetTable): %w", err)
		}
		users = append(users, board)
	}
	return users, nil
}

func (r *BoardRepository) FindItemByID(id int) (Board, error) {
	row := r.db.QueryRow("SELECT * FROM boards WHERE id = $1", id)
	if row.Err() != nil {
		return Board{}, fmt.Errorf("In BoardRepository(FindItemByID): %w", row.Err())
	}
	var board Board
	err := row.Scan(&board.ID, &board.OwnerID, &board.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNotFound
		}
		err = fmt.Errorf("In BoardRepository(FindItemByID): %w", err)
	}
	return board, err
}

func (r *BoardRepository) FindItemByCondition(condition func(item Board) bool) (Board, error) {
	items, err := r.FindItemsByCondition(condition)
	if err != nil {
		return Board{}, fmt.Errorf("In BoardRepository(FindItemByCondition): %w", err)
	}
	if len(items) == 0 {
		return Board{}, ErrNotFound
	}
	return items[0], nil
}

func (r *BoardRepository) FindItemsByCondition(condition func(item Board) bool) ([]Board, error) {
	table, err := r.GetTable()
	if err != nil {
		return nil, fmt.Errorf("In BoardRepository(FindItemsByCondition): %w", err)
	}
	res := make([]Board, 0) //???
	for _, board := range table {
		if condition(board) {
			res = append(res, board)
		}
	}
	return res, nil
}
