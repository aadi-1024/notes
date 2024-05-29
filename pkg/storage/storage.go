package storage

import "github.com/aadi-1024/notes/pkg/models"

type Storage interface {
	RegisterUser(models.User) error
	LoginUser(models.User) ([]byte, error)

	CreateNote(note models.Note)
	DeleteNote(note models.Note)
	UpdateNote(note models.Note)
	GetNote(note models.Note)
	GetAll(userId int)
	//GetN returns all notes on page number page, if they are divided as n per page
	GetN(userId, page, n int)
}
