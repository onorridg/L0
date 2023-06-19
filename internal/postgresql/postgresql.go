package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"l0/internal/env"
	"l0/internal/models"
	"l0/internal/parser"
	"log"
)

type transaction struct {
	tx    *sql.Tx
	txCtx context.Context
}

type DB struct {
	Conn *sql.DB
	transaction
}

func (db *DB) insertItems(order *models.Order) {
	var err error

	for i := range order.Items {
		queryStr := `
		INSERT INTO item (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
		_, err = db.tx.ExecContext(db.txCtx, queryStr, parser.ItemStructToSlice(i, order))
		if err != nil {
			db.tx.Rollback()
			log.Println(err)
			return
		}
	}
}

func (db *DB) InsertTransaction(order *models.Order) {
	var err error
	log.Println("[!] Start transaction")
	db.txCtx = context.Background()
	db.tx, err = db.Conn.BeginTx(db.txCtx, nil)
	if err != nil {
		log.Fatal("Init tx:", err)
	}

	queryStr := `
	INSERT INTO user_order (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, name, phone, zip, city, address, region, email, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27)`
	_, err = db.tx.ExecContext(db.txCtx, queryStr, parser.OrderStructToSlice(order))
	if err != nil {
		db.tx.Rollback()
		log.Println(err)
		return
	}
	//db.insertItems(order)

	err = db.tx.Commit()
	if err != nil {
		//log.Fatal(err)
		log.Println("commit:", err)
	}
	log.Println("Commit success!")
}

func (db *DB) PushInMemory() {

}

func initConn() *sql.DB {
	user := env.Get().PgUser
	password := env.Get().PgPassword
	host := env.Get().PgHost
	port := env.Get().PgPort
	dbName := env.Get().PgDatabase
	connStr := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable", user, password, host, port, dbName)
	// Устанавливаем соединение с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	// Проверяем подключение
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
