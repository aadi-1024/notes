package models

import "encoding/gob"

type Session struct {
	User *User
}

func init() {
	gob.Register(User{})
	gob.Register(Note{})
	gob.Register(Session{})
}
