package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"` // for password reset
	Password string `json:"password"`
}
