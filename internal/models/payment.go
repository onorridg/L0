package models

import (
	t "l0/internal/models/castomType"
)

type Payment struct {
	Transaction  t.FlexStr   `json:"transaction" db:"transaction"`
	RequestID    t.FlexStr   `json:"request_id" db:"request_id"`
	Currency     string      `json:"currency" db:"currency"`
	Provider     string      `json:"provider" db:"provider"`
	Amount       t.FlexFloat `json:"amount" db:"amount"`
	PaymentDt    t.FlexStr   `json:"payment_dt" db:"payment_dt"`
	Bank         string      `json:"bank" db:"bank"`
	DeliveryCost t.FlexFloat `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   t.FlexFloat `json:"goods_total" db:"goods_total"`
	CustomFee    t.FlexStr   `json:"custom_fee" db:"custom_fee"`
}
