package tests

import (
	"log"
	"testing"
	"time"

	"github.com/aadi-1024/notes/pkg/database"
	"github.com/aadi-1024/notes/pkg/models"
)

var db *database.Database

func init() {
	d, err := database.InitDatabase("postgres://postgres:password@localhost:5432/notes", 3*time.Second)
	if err != nil {
		panic(err.Error())
	}
	db = d
	log.Println()
}

func TestRegisterUser(t *testing.T) {
	u := models.User{
		Username: "user",
		Email:    "user@email.com",
		Password: "password",
	}
	err := db.RegisterUser(u)
	t.Log("here")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestLoginUser(t *testing.T) {
	u := models.User{
		Username: "user",
		Email:    "user@email.com",
		Password: "password",
	}
	_, err := db.LoginUser(u)
	if err != nil {
		t.Error(err.Error())
	}
}
