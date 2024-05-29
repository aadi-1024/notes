package database

import (
	"context"
	"log"
	"time"

	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool    *pgxpool.Pool
	timeout time.Duration
}

func InitDatabase(dsn string, timeout time.Duration) (*Database, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 5; i++ {
		log.Println("pinging database")
		err = pool.Ping(context.Background())
		if err != nil {
			log.Printf("ping %v failed", i)
			time.Sleep(timeout)
		} else {
			log.Println("successfully connected")
			break
		}
	}
	// would be nil if ping succeeded
	if err != nil {
		log.Println("could not connect to database, exiting")
		return nil, err
	}

	db := &Database{
		pool:    pool,
		timeout: timeout,
	}
	return db, nil
}
