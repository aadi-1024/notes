package models

import "time"

type Note struct {
	Id        int       `json:"id" validate:"gte=1"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title" validate:"required,gte=1,lte=256"`
	Text      string    `json:"text" validate:"required,gte=0,lte=2048"`
	CreatedAt time.Time `json:"created_at"`
}
