package models

type User struct {
	ID       uint   `json:"id"`
	Login    string `json:"login"`
	Password []byte `json:"password"`
}
