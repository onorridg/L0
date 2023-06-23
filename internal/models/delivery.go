package models

import (
	t "l0/internal/models/castomType"
)

type Delivery struct {
	Name    string    `json:"name" db:"name"`
	Phone   string    `json:"phone" db:"phone"`
	Zip     t.FlexStr `json:"zip" db:"zip"`
	City    string    `json:"city" db:"city"`
	Address string    `json:"address" db:"address"`
	Region  string    `json:"region" db:"region"`
	Email   string    `json:"email" db:"email"`
}
