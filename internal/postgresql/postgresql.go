package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
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

type BadRequest struct {
	Err string
}

func (db *DB) insertItems(order *models.Order, userOrderId int64) error {
	var err error

	for i := range order.Items {
		queryStr := `
		INSERT INTO item (user_order_id, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
		_, err = db.tx.ExecContext(db.txCtx, queryStr, parser.ItemStructToSlice(i, order, userOrderId)...)
		if err != nil {
			db.tx.Rollback()
			return err
		}
	}
	return nil
}

func (db *DB) insertOrder(order *models.Order) (int64, error) {
	queryStr := `
	INSERT INTO user_order (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard, name, phone, zip, city, address, region, email, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28)
	RETURNING id`

	var id int64
	err := db.tx.QueryRowContext(db.txCtx, queryStr, parser.OrderStructToSlice(order)...).Scan(&id)
	if err != nil {
		db.tx.Rollback()
		return 0, err
	}
	return id, nil
}

func (db *DB) InsertUserOrder(order *models.Order) {
	var err error

	log.Println("[+] Start transaction.")
	db.txCtx = context.Background()
	db.tx, err = db.Conn.BeginTx(db.txCtx, nil)
	if err != nil {
		log.Println("[!] Init tx:", err)
		return
	}

	userOrderId, err := db.insertOrder(order)
	if err != nil {
		log.Println("[!] Commit fail.\n", err)
		return
	}

	err = db.insertItems(order, userOrderId)
	if err != nil {
		log.Println("[!] Commit fail.\n", err)
		return
	}

	err = db.tx.Commit()
	if err != nil {
		log.Println("[!] Commit fail.\n", err)
		return
	}

	log.Println("[+] Commit success. ID:", userOrderId)
}

func (db *DB) selectItems(orderId uint64, dbx *sqlx.DB) []models.Items {
	queryStr := `SELECT * FROM item WHERE user_order_id = $1`
	var userItems []models.Items

	err := dbx.Select(&userItems, queryStr, orderId)
	if err != nil {
		log.Println(err)
	}
	return userItems
}

func (db *DB) SelectUsrOrder(orderId uint64) (uint64, any, error) {
	dbx := sqlx.NewDb(db.Conn, "postgres")

	queryStr := `SELECT * FROM user_order WHERE id = $1`
	var userOrder models.Order

	err := dbx.Get(&userOrder, queryStr, orderId)
	if err != nil {
		return 0, &BadRequest{Err: err.Error()}, err
	}

	userOrder.Items = db.selectItems(orderId, dbx)

	return userOrder.Id, &userOrder, nil
}

func initConn() *sql.DB {
	user := env.Get().PgUser
	password := env.Get().PgPassword
	host := env.Get().PgHost
	port := env.Get().PgPort
	dbName := env.Get().PgDatabase
	connStr := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable", user, password, host, port, dbName)

	db, err := sql.Open("postgres", connStr)
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
