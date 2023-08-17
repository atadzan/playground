package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func connect() (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf("postgres://tmtube_admin:3Qv@e8U0ImT@localhost:5454/tmtube?sslmode=disable")
	return pgxpool.Connect(context.Background(), dbURL)
}

func batchInsert(db *pgxpool.Pool) error {
	batch := &pgx.Batch{}
	for i := 0; i < 100; i++ {
		batch.Queue("INSERT INTO test_table(title) values $1", fmt.Sprintf("Hello-%d", i))
	}
	_ = db.SendBatch(context.Background(), batch)
	return nil
}

func main() {
	db, err := connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	if err = batchInsert(db); err != nil {
		log.Println(err)
		return
	}

}
