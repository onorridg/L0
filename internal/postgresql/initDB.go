package postgresql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"l0/internal/env"
)

type Transaction struct {
	tx *sqlx.Tx
}

type DB struct {
	Conn *sqlx.DB
	Transaction
}

func initConn() *sqlx.DB {
	user := env.Get().PgUserWorker
	password := env.Get().PgPasswordWorker
	host := env.Get().PgHost
	port := env.Get().PgPort
	dbName := env.Get().PgDatabase
	connStr := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable", user, password, host, port, dbName)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func Conn() *DB {
	newDb := DB{}
	newDb.Conn = initConn()
	return &newDb
}
