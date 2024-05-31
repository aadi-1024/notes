package storage

import "github.com/aadi-1024/notes/pkg/models"

type Storage interface {
	RegisterUser(models.User) error
	LoginUser(models.User) ([]byte, error)

	CreateNote(note models.Note) (int, error)
	DeleteNote(note models.Note) error
	UpdateNote(note models.Note) error
	GetNote(note models.Note) (models.Note, error)
	GetAll(userId int) ([]models.Note, error)
	//GetN returns all notes on page number page, if they are divided as n per page
	GetN(userId, page, n int) ([]models.Note, error)
}
