package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type FlexFloat float64

func (ff *FlexFloat) Scan(value interface{}) error {
	if value == nil {
		*ff = 0.0
		return nil
	}

	if v, ok := value.(int64); ok {
		*ff = FlexFloat(v)
		return nil
	} else if v, ok := value.(float64); ok {
		*ff = FlexFloat(v)
		return nil
	} else if v, ok := value.(string); ok {
		res, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		*ff = FlexFloat(res)
		return nil
	}

	return fmt.Errorf("failed to scan FlexStr value")
}

func (ff *FlexFloat) UnmarshalJSON(b []byte) error {
	var value interface{}
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	switch v := value.(type) {
	case float64:
		*ff = FlexFloat(v)
	case string:
		fNum, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		*ff = FlexFloat(fNum)
	default:
		return fmt.Errorf("unsupported value type: %T", v)
	}
	return nil

}

type FlexStr string

func (fs *FlexStr) Scan(value interface{}) error {
	if value == nil {
		*fs = ""
		return nil
	}

	if v, ok := value.(int64); ok {
		*fs = FlexStr(strconv.FormatInt(v, 10))
		return nil
	} else if v, ok := value.(float64); ok {
		*fs = FlexStr(strconv.FormatFloat(v, 'f', -1, 64))
		return nil
	} else if v, ok := value.(string); ok {
		*fs = FlexStr(v)
		return nil
	}

	return fmt.Errorf("failed to scan FlexStr value")
}

func (fs *FlexStr) UnmarshalJSON(b []byte) error {
	var value interface{}
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	switch v := value.(type) {
	case float64:
		*fs = FlexStr(strconv.FormatInt(int64(v), 10))
	case string:
		*fs = FlexStr(v)
	default:
		return fmt.Errorf("unsupported value type: %T", v)
	}
	return nil
}

type Delivery struct {
	Name    string  `json:"name" db:"name"`
	Phone   string  `json:"phone" db:"phone"`
	Zip     FlexStr `json:"zip" db:"zip"`
	City    string  `json:"city" db:"city"`
	Address string  `json:"address" db:"address"`
	Region  string  `json:"region" db:"region"`
	Email   string  `json:"email" db:"email"`
}

type Payment struct {
	Transaction  FlexStr   `json:"transaction" db:"transaction"`
	RequestID    FlexStr   `json:"request_id" db:"request_id"`
	Currency     string    `json:"currency" db:"currency"`
	Provider     string    `json:"provider" db:"provider"`
	Amount       FlexFloat `json:"amount" db:"amount"`
	PaymentDt    FlexStr   `json:"payment_dt" db:"payment_dt"`
	Bank         string    `json:"bank" db:"bank"`
	DeliveryCost FlexFloat `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   FlexFloat `json:"goods_total" db:"goods_total"`
	CustomFee    FlexStr   `json:"custom_fee" db:"custom_fee"`
}

type Items struct {
	Id          uint64    `db:"id"`
	UserOrderId uint64    `db:"user_order_id"`
	ChrtID      FlexStr   `json:"chrt_id" db:"chrt_id"`
	TrackNumber string    `json:"track_number" db:"track_number"`
	Price       FlexFloat `json:"price" db:"price"`
	Rid         string    `json:"rid" db:"rid"`
	Name        string    `json:"name" db:"name"`
	Sale        FlexFloat `json:"sale" db:"sale"`
	Size        FlexFloat `json:"size" db:"size"`
	TotalPrice  FlexFloat `json:"total_price" db:"total_price"`
	NmID        FlexStr   `json:"nm_id" db:"nm_id"`
	Brand       string    `json:"brand" db:"brand"`
	Status      FlexStr   `json:"status" db:"status"`
}

type Order struct {
	Id                uint64 `db:"id"`
	OrderUID          string `json:"order_uid" db:"order_uid"`
	TrackNumber       string `json:"track_number" db:"track_number"`
	Entry             string `json:"entry" db:"entry"`
	Delivery          `json:"delivery"`
	Payment           `json:"payment"`
	Items             []Items `json:"items"`
	Locale            string  `json:"locale" db:"locale"`
	InternalSignature string  `json:"internal_signature" db:"internal_signature"`
	CustomerID        string  `json:"customer_id" db:"customer_id"`
	DeliveryService   string  `json:"delivery_service" db:"delivery_service"`
	Shardkey          FlexStr `json:"shardkey" db:"shardkey"`
	SmID              FlexStr `json:"sm_id" db:"sm_id"`
	DateCreated       string  `json:"date_created" db:"date_created"`
	OofShard          FlexStr `json:"oof_shard" db:"oof_shard"`
}

type kek struct {
	ms int
}

type lol struct {
	id int
	kek
}
