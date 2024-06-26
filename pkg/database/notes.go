package database

import (
	"context"
	"errors"
	"time"

	"github.com/aadi-1024/notes/pkg/models"
	"github.com/jackc/pgx/v5"
)

func (d *Database) CreateNote(note models.Note) (int, error) {
	query := `INSERT INTO notes (UserId, Title, Text, CreatedAt) VALUES ($1, $2, $3, $4) RETURNING Id;`

	tx, err := d.pool.BeginTx(context.Background(), pgx.TxOptions{})
	defer tx.Rollback(context.Background())
	if err != nil {
		return 0, err
	}

	row := tx.QueryRow(context.Background(), query, note.UserId, note.Title, note.Text, time.Now())
	var id int
	err = row.Scan(&id)
	if err == nil {
		return id, tx.Commit(context.Background())
	}
	return 0, err
}

func (d *Database) DeleteNote(id, uid int) error {
	query := `delete from notes where Id = $1 and UserId = $2;`

	tx, err := d.pool.BeginTx(context.Background(), pgx.TxOptions{})
	defer tx.Rollback(context.Background())
	if err != nil {
		return err
	}

	res, err := tx.Exec(context.Background(), query, id, uid)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("not found")
	}
	return tx.Commit(context.Background())
}

func (d *Database) UpdateNote(note models.Note) error {
	query := `update notes set Title = $1, Text = $2 where Id = $3 and UserId = $4;`

	tx, err := d.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), query, note.Title, note.Text, note.Id, note.UserId)
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (d *Database) GetNote(note models.Note) (models.Note, error) {
	query := `select * from notes where Id = $1 and UserId = $2;`

	row := d.pool.QueryRow(context.Background(), query, note.Id, note.UserId)
	ret := models.Note{}

	err := row.Scan(&ret.Id, &ret.UserId, &ret.Title, &ret.Text, &ret.CreatedAt)
	return ret, err
}

func (d *Database) GetAll(userId int) ([]models.Note, error) {
	query := `select * from notes where UserId = $1;`

	rows, err := d.pool.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Note, error) {
		note := models.Note{}
		err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Text, &note.CreatedAt)
		return note, err
	})

	return notes, err
}

func (d *Database) GetN(userId, page, n int) ([]models.Note, error) {
	query := `select * from notes where UserId = $1 order by CreatedAt desc offset $2 rows limit $3`

	rows, err := d.pool.Query(context.Background(), query, userId, (page-1)*n, n)
	if err != nil {
		return nil, err
	}

	rows.Close()

	notes, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Note, error) {
		note := models.Note{}
		err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Text, &note.CreatedAt)
		return note, err
	})
	return notes, err
}
