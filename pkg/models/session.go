package models

import "encoding/gob"

type Session struct {
	LoggedIn  bool
	SessionId string
	User      *User
	Notes     []*Note
}

func init() {
	gob.Register(User{})
	gob.Register(Note{})
	gob.Register(Session{})
}
