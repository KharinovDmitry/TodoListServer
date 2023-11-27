package models

type Column struct {
	ID      uint   `json:"id"`
	BoardID uint   `json:"boardID"`
	Name    string `json:"name"`
}
