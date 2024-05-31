package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"gte=6,lte=24"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=24"`
}
