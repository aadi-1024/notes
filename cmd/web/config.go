package main

import (
	"github.com/aadi-1024/notes/pkg/database"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	debug     bool
	db        *database.Database
	session   *scs.SessionManager
	validator *validator.Validate
}
