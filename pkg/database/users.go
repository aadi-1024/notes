package database

import (
	"context"

	"github.com/aadi-1024/notes/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *Database) RegisterUser(user models.User) error {
	query := `INSERT INTO users (Username, Email, Password) VALUES ($1, $2, $3) RETURNING Id`

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), -1)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()
	_, err = d.pool.Exec(ctx, query, user.Username, user.Email, hash)
	return err
}

// LoginUser returns a session token, for now just returning a dummy value
func (d *Database) LoginUser(user models.User) (int, error) {
	query := `SELECT * FROM users WHERE Email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()

	ret := models.User{}
	row := d.pool.QueryRow(ctx, query, user.Email)
	err := row.Scan(&ret.Id, &ret.Username, &ret.Email, &ret.Password)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(ret.Password), []byte(user.Password))
	if err != nil {
		return 0, err
	}
	return ret.Id, nil
}
