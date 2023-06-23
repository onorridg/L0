package postgresql

import (
	"l0/internal/models"
	"log"
)

func (db *DB) insertItems(order *models.Order, userOrderId uint64) error {
	var err error

	for i := range order.Items {
		order.Items[i].UserOrderId = userOrderId
		queryStr := `
		INSERT INTO item (user_order_id, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
		VALUES (:user_order_id, :chrt_id, :track_number, :price, :rid, :name, :sale, :size, :total_price, :nm_id, :brand, :status)
		`
		_, err = db.tx.NamedExec(queryStr, order.Items[i])
		if err != nil {
			log.Println(i, "-", order.Items[i])
			db.tx.Rollback()
			return err
		}

	}
	return err
}

func (db *DB) insertOrder(order *models.Order) (uint64, error) {
	queryStr := `
	INSERT INTO user_order (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard, name, phone, zip, city, address, region, email, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
	VALUES (:order_uid, :track_number, :entry, :locale, :internal_signature, :customer_id, :delivery_service, :shardkey, :sm_id, :date_created, :oof_shard, :name, :phone, :zip, :city, :address, :region, :email, :transaction, :request_id, :currency, :provider, :amount, :payment_dt, :bank, :delivery_cost, :goods_total, :custom_fee)
	RETURNING id`

	var id uint64
	rows, err := db.tx.NamedQuery(queryStr, order)
	if err != nil {
		db.tx.Rollback()
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			db.tx.Rollback()
			return 0, err
		}
	}
	return id, nil
}

func (db *DB) InsertUserOrder(order *models.Order) {
	var err error

	log.Println("[+] Start transaction.")
	db.tx, err = db.Conn.Beginx()
	if err != nil {
		log.Println("[!] Init tx:", err)
		return
	}

	userOrderId, err := db.insertOrder(order)
	if err != nil {
		log.Println("[!] Commit fail (insert order).", err)
		return
	}

	err = db.insertItems(order, userOrderId)
	if err != nil {
		log.Println("[!] Commit fail (insert item).", err)
		return
	}

	err = db.tx.Commit()
	if err != nil {
		log.Println("[!] Commit fail (commit).", err)
		return
	}

	log.Println("[+] Commit success. ID:", userOrderId)
}
