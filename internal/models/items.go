package models

import (
	t "l0/internal/models/castomType"
)

type Items struct {
	Id          uint64      `json:"id" db:"id"`
	UserOrderId uint64      `db:"user_order_id"`
	ChrtID      t.FlexStr   `json:"chrt_id" db:"chrt_id"`
	TrackNumber string      `json:"track_number" db:"track_number"`
	Price       t.FlexFloat `json:"price" db:"price"`
	Rid         string      `json:"rid" db:"rid"`
	Name        string      `json:"name" db:"name"`
	Sale        t.FlexFloat `json:"sale" db:"sale"`
	Size        t.FlexFloat `json:"size" db:"size"`
	TotalPrice  t.FlexFloat `json:"total_price" db:"total_price"`
	NmID        t.FlexStr   `json:"nm_id" db:"nm_id"`
	Brand       string      `json:"brand" db:"brand"`
	Status      t.FlexStr   `json:"status" db:"status"`
}
