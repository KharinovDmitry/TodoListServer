package postgres

import (
	. "TodoListServer/internal/domain/models"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) AddItem(item User) error {
	_, err := r.db.Exec("INSERT INTO users(login, password) VALUES ($1, $2)",
		item.Login, item.Password)
	if err != nil {
		err = fmt.Errorf("In UserRepository(AddItem): %w", err)
	}
	return err
}

func (r *UserRepository) DeleteItem(item User) error {
	panic("not implemented")
}

func (r *UserRepository) UpdateItem(id int, newItem User) error {
	panic("not implemented")
}

func (r *UserRepository) GetTable() ([]User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("In UserRepository(GetTable): %w", err)
	}
	defer rows.Close()

	users := make([]User, 0) //??
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("In UserRepository(GetTable): %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) FindItemByCondition(condition func(item User) bool) (User, error) {
	items, err := r.FindItemsByCondition(condition)
	if err != nil {
		return User{}, fmt.Errorf("In UserRepository(FindItemByCondition): %w", err)
	}
	if len(items) == 0 {
		return User{}, ErrNotFound
	}
	return items[0], nil
}

func (r *UserRepository) FindItemsByCondition(condition func(item User) bool) ([]User, error) {
	table, err := r.GetTable()
	if err != nil {
		return nil, fmt.Errorf("In UserRepository(FindItemsByCondition): %w", err)
	}
	res := make([]User, 0) //???
	for _, user := range table {
		if condition(user) {
			res = append(res, user)
		}
	}
	return res, nil
}

func (r *UserRepository) FindItemByID(id int) (User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	if row.Err() != nil {
		return User{}, fmt.Errorf("In UserRepository(FindItemByID): %w", row.Err())
	}
	var user User
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNotFound
		}
		err = fmt.Errorf("In UserRepository(FindItemByID): %w", err)
	}
	return user, err
}
