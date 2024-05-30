package main

import (
	"github.com/aadi-1024/notes/pkg/database"
	"github.com/alexedwards/scs/v2"
)

type Config struct {
	debug   bool
	db      *database.Database
	session *scs.SessionManager
}
