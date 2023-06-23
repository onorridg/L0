package postgresql

import (
	"fmt"
	"l0/internal/env"
	"l0/internal/models"
	"log"
)

func setItemsToOrders(orders []models.Order, items []models.Items) []models.Order {
	itemsLen := len(items)
	itemIndex := 0

	for orderIndex := range orders {
		for ; itemIndex < itemsLen; itemIndex++ {
			if orders[orderIndex].Id != items[itemIndex].UserOrderId {
				break
			}
			orders[orderIndex].Items = append(orders[orderIndex].Items, items[itemIndex])
		}
	}

	return orders
}

func (db *DB) getAllItemsForNOrders(startId uint64) []models.Items {
	var items []models.Items

	queryStr := fmt.Sprint("SELECT * FROM item WHERE user_order_id >=", startId)
	err := db.tx.Select(&items, queryStr)
	if err != nil {
		log.Fatal("[!] Commit fail (select orders).", err)
	}

	return items
}

func (db *DB) getNOrders() []models.Order {
	var orders []models.Order

	queryStr := fmt.Sprint("SELECT * FROM user_order LIMIT", env.Get().CacheSize)
	err := db.tx.Select(&orders, queryStr)
	if err != nil {
		log.Fatal("[!] Commit fail (select orders).", err)
	}

	return orders
}

func (db *DB) GetLastNOrders() []models.Order {
	var err error

	db.tx, err = db.Conn.Beginx()
	if err != nil {
		log.Fatal("[!] Init tx:", err)
	}

	orders := db.getNOrders()
	if len(orders) != 0 {
		items := db.getAllItemsForNOrders(orders[0].Id)
		orders = setItemsToOrders(orders, items)
	}

	err = db.tx.Commit()
	if err != nil {
		log.Fatal("[!] Commit fail (commit).", err)
	}

	if len(orders) == 0 {
		return nil
	}
	return orders
}
