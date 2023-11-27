package models

type Task struct {
	ID          uint   `json:"id"`
	ColumnID    uint   `json:"columnID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
