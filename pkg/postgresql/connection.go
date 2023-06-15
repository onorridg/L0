package postgresql

import (
	"database/sql"
	"fmt"
	"l0/internal/env"
)

type db struct {
	Conn *sql.DB
}

func (db *db) Insert() {

}

func (db *db) PushInMemory() {

}

func initConn() *sql.DB {
	user := env.Get().PgUser
	password := env.Get().PgPassword
	host := env.Get().PgHost
	port := env.Get().PgPort
	dbName := env.Get().PgDatabase
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbName)

	// Устанавливаем соединение с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func Conn() *db {
	newDb := db{}
	newDb.Conn = initConn()
	return &newDb
}
