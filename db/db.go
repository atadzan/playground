package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"strings"
	"time"
)

func connect(ctx context.Context) (*pgx.Conn, error) {
	dbURL := fmt.Sprintf("postgres://tmtube_admin:3Qv@e8U0ImT@localhost:5454/tmtube?sslmode=disable")
	return pgx.Connect(ctx, dbURL)
}

/* INSERT FUNCTIONS */
func batchInsert(ctx context.Context, db *pgx.Conn) error {
	batch := &pgx.Batch{}

	for i := 0; i < 1000; i++ {
		batch.Queue("INSERT INTO test_table(title) values($1)", fmt.Sprintf("Hello-%d", i))
	}
	startTime := time.Now()
	br := db.SendBatch(ctx, batch)
	defer br.Close()
	_, err := br.Exec()
	if err != nil {
		log.Println(err)
	}
	log.Println("Inserts via pg copy protocol elapsed time:", time.Since(startTime).Milliseconds(), "ms")
	return nil
}

func insertByOne(ctx context.Context, db *pgx.Conn) error {
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		_, err := db.Exec(ctx, "INSERT INTO test_table(title) VALUES($1)", fmt.Sprintf("object-%d", i))
		if err != nil {
			log.Println(err)
			return err
		}
	}

	log.Println("Inserts via insert by one elapsed time:", time.Since(startTime).Milliseconds(), "ms")
	return nil
}

func insertMultiValues(ctx context.Context, db *pgx.Conn) error {
	var rawValues []string
	for i := 0; i < 1000; i++ {
		rawValues = append(rawValues, fmt.Sprintf("('object-%d')", i))
	}
	values := strings.Join(rawValues, ",")
	startTime := time.Now()

	_, err := db.Exec(ctx, fmt.Sprintf("INSERT INTO test_table(title) VALUES %s", values))
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Inserts via multiple values elapsed time:", time.Since(startTime).Milliseconds(), "ms")
	return nil
}

func insertByTx(ctx context.Context, db *pgx.Conn) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Commit(ctx)
	var rawValues []string
	for i := 0; i < 1000; i++ {
		rawValues = append(rawValues, fmt.Sprintf("('object-%d')", i))
	}
	values := strings.Join(rawValues, ",")
	startTime := time.Now()

	_, err = tx.Exec(ctx, fmt.Sprintf("INSERT INTO test_table(title) VALUES %s", values))
	if err != nil {
		log.Println(err)
		_ = tx.Rollback(ctx)
		return err
	}

	log.Println("Inserts via multiple values using TX elapsed time:", time.Since(startTime).Milliseconds(), "ms")

	return nil
}

func insertCopyFrom(ctx context.Context, db *pgx.Conn) error {
	sources := [][]interface{}{}
	for i := 0; i < 1000; i++ {
		sources = append(sources, []interface{}{fmt.Sprintf("object-%d", i)})
	}
	startTime := time.Now()
	_, err := db.CopyFrom(ctx, pgx.Identifier{"test_table"}, []string{"title"}, pgx.CopyFromRows(sources))
	if err != nil {
		log.Println("Error while inserting", err)
		return err
	}

	log.Println("Inserts via pg copy protocol elapsed time:", time.Since(startTime).Milliseconds(), "ms")
	return nil
}

/* UPDATE FUNCTIONS */
func batchUpdate(ctx context.Context, db *pgx.Conn) error {
	batch := &pgx.Batch{}

	for i := 0; i < 1000; i++ {
		batch.Queue("UPDATE test_table SET title=$1", fmt.Sprintf("batch-%d", i))
	}
	startTime := time.Now()
	br := db.SendBatch(ctx, batch)
	defer br.Close()
	_, err := br.Exec()
	if err != nil {
		log.Println(err)
	}
	log.Println("Update via batch elapsed time:", time.Since(startTime).Milliseconds(), "ms")
	return nil
}

func updateByOne(ctx context.Context, db *pgx.Conn) error {
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		_, err := db.Exec(ctx, "UPDATE test_table SET title=$1", fmt.Sprintf("updateByOne-%d", i))
		if err != nil {
			log.Println(err)
			return err
		}
	}

	log.Println("Update via update by one elapsed time:", time.Since(startTime).Milliseconds(), "ms")
	return nil
}

func updateMultiValues(ctx context.Context, db *pgx.Conn) error {
	var rawValues []string
	for i := 0; i < 5; i++ {
		rawValues = append(rawValues, fmt.Sprintf("('updateMulti-%d')", i))
	}
	values := strings.Join(rawValues, ",")
	startTime := time.Now()

	_, err := db.Exec(ctx, fmt.Sprintf("UPDATE test_table SET title=c.title FROM (values(%s)) as c(title)", values))
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Update via multiple values elapsed time:", time.Since(startTime).Milliseconds(), "ms")
	return nil
}

func updateByTx(ctx context.Context, db *pgx.Conn) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Commit(ctx)
	var rawValues []string
	for i := 0; i < 5; i++ {
		rawValues = append(rawValues, fmt.Sprintf("('updateTX-%d')", i))
	}
	values := strings.Join(rawValues, ",")
	startTime := time.Now()

	_, err = db.Exec(ctx, fmt.Sprintf("UPDATE test_table SET title=c.title FROM (values(%s)) as c(title)", values))
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Update via multiple values using TX elapsed time:", time.Since(startTime).Milliseconds(), "ms")

	return nil
}

func insertFunctions(ctx context.Context, db *pgx.Conn) {
	log.Println("Total object which will be inserted 1000")
	/* Insert by one */
	/*
		Example
		INSERT INTO test_table VALUES('object-1')
		INSERT INTO test_table VALUES('object-2')

	*/
	if err := insertByOne(ctx, db); err != nil {
		log.Println(err)
		return
	}

	//* Insert with multiple values  */
	/*
		Example query
		INSERT INTO test_table VALUES('object-1'), ('object-2')
	*/
	if err := insertMultiValues(ctx, db); err != nil {
		log.Println(err)
		return
	}

	/* Insert via postgres COPY protocol */

	if err := insertCopyFrom(ctx, db); err != nil {
		log.Println(err)
		return
	}

	//* Insert via batch  */
	/*
		Example query
	*/
	if err := batchInsert(ctx, db); err != nil {
		log.Println(err)
		return
	}
	//* Insert via tx  */
	/*
		Example query
	*/
	if err := insertByTx(ctx, db); err != nil {
		log.Println(err)
		return
	}
}

func updateFunctions(ctx context.Context, db *pgx.Conn) {
	log.Println("Total object which will be updated 1000")
	//* Update via update by one  */
	if err := updateByOne(ctx, db); err != nil {
		log.Println(err)
		return
	}

	//* Update via batch  */
	if err := batchUpdate(ctx, db); err != nil {
		log.Println(err)
		return
	}

	//* Update via updateMultiValues  */
	/*	Example query
		update test as t set
			column_a = c.column_a
			from (values
			('123', 1),
			('345', 2)
		) as c(column_b, column_a)
			where c.column_b = t.column_b;
	*/
	if err := updateMultiValues(ctx, db); err != nil {
		log.Println(err)
		return
	}

	/* Update via update by TX  */
	if err := updateByTx(ctx, db); err != nil {
		log.Println(err)
		return
	}
}

func clearTable(ctx context.Context, db *pgx.Conn) {
	query := fmt.Sprintf(`DELETE FROM test_table`)
	_, err := db.Exec(ctx, query)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table is cleaned")
}

func main() {
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println(err)
	}
	defer db.Close(ctx)

	//insertFunctions(ctx, db)

	updateFunctions(ctx, db)

	//clearTable(ctx, db)
}
