package models

import "encoding/gob"

type Session struct {
	LoggedIn bool
	UserId   int
	User     *User
}

func init() {
	gob.Register(User{})
	gob.Register(Note{})
	gob.Register(Session{})
}
