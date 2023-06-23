package models

import (
	t "l0/internal/models/castomType"
)

type Order struct {
	Id                uint64 `json:"id" db:"id"`
	OrderUID          string `json:"order_uid" db:"order_uid"`
	TrackNumber       string `json:"track_number" db:"track_number"`
	Entry             string `json:"entry" db:"entry"`
	Delivery          `json:"delivery"`
	Payment           `json:"payment"`
	Items             []Items   `json:"items"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerID        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	Shardkey          t.FlexStr `json:"shardkey" db:"shardkey"`
	SmID              t.FlexStr `json:"sm_id" db:"sm_id"`
	DateCreated       string    `json:"date_created" db:"date_created"`
	OofShard          t.FlexStr `json:"oof_shard" db:"oof_shard"`
}
