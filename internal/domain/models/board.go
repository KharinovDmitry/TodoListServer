package models

type Board struct {
	ID      uint   `json:"id"`
	OwnerID uint   `json:"ownerID"`
	Name    string `json:"name"`
}
