package postgresql

import (
	"l0/internal/models"
	"log"
)

type badRequest struct {
	Err string
}

func (db *DB) selectItems(orderId uint64) []models.Items {
	queryStr := `SELECT * FROM item WHERE user_order_id = $1`
	var userItems []models.Items

	err := db.Conn.Select(&userItems, queryStr, orderId)
	if err != nil {
		log.Println(err)
	}
	return userItems
}

func (db *DB) SelectUserOrder(orderId uint64) (uint64, any, error) {
	queryStr := `SELECT * FROM user_order WHERE id = $1`
	var userOrder models.Order

	// Get order
	err := db.Conn.Get(&userOrder, queryStr, orderId)
	if err != nil {
		return 0, &badRequest{Err: err.Error()}, err
	}

	// Set items for order
	userOrder.Items = db.selectItems(orderId)

	return userOrder.Id, &userOrder, nil
}
