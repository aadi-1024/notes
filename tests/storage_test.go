package tests

import (
	"log"
	"testing"
	"time"

	"github.com/aadi-1024/notes/pkg/database"
	"github.com/aadi-1024/notes/pkg/models"
)

var noteId int
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

func TestCreateNote(t *testing.T) {
	n := models.Note{
		UserId:    1,
		Title:     "New Note",
		Text:      "New Note Text",
		CreatedAt: time.Now(),
	}

	id, err := db.CreateNote(n)
	if err != nil {
		t.Fatal(err.Error())
	}
	noteId = id
}

func TestGetNote(t *testing.T) {
	n := models.Note{
		Id:     noteId,
		UserId: 1,
	}
	ret, err := db.GetNote(n)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(ret)
}

func TestGetAll(t *testing.T) {
	userId := 1
	db.CreateNote(models.Note{
		UserId:    userId,
		Title:     "Another Note",
		Text:      "Another Note Text",
		CreatedAt: time.Now(),
	})
	notes, err := db.GetAll(userId)
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range notes {
		t.Log(v)
	}
}

func TestGetN(t *testing.T) {
	userId := 1
	_, err := db.CreateNote(models.Note{
		UserId: userId,
		Title:  "Wow Note",
		Text:   "Wow note text",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	notes, err := db.GetN(1, 2, 2)
	if err != nil || len(notes) != 2 {
		t.Fatal(err.Error())
	}
	for _, v := range notes {
		t.Log(v)
	}
}
