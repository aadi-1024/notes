package main

import "github.com/aadi-1024/notes/pkg/database"

type Config struct {
	debug bool
	db    *database.Database
}
