package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type FlexStr string

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

type Order struct {
	OrderUID    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string  `json:"name"`
		Phone   string  `json:"phone"`
		Zip     FlexStr `json:"zip"`
		City    string  `json:"city"`
		Address string  `json:"address"`
		Region  string  `json:"region"`
		Email   string  `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  FlexStr `json:"transaction"`
		RequestID    FlexStr `json:"request_id"`
		Currency     string  `json:"currency"`
		Provider     string  `json:"provider"`
		Amount       FlexStr `json:"amount"`
		PaymentDt    FlexStr `json:"payment_dt"`
		Bank         string  `json:"bank"`
		DeliveryCost FlexStr `json:"delivery_cost"`
		GoodsTotal   FlexStr `json:"goods_total"`
		CustomFee    FlexStr `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtID      FlexStr `json:"chrt_id"`
		TrackNumber string  `json:"track_number"`
		Price       FlexStr `json:"price"`
		Rid         string  `json:"rid"`
		Name        string  `json:"name"`
		Sale        FlexStr `json:"sale"`
		Size        FlexStr `json:"size"`
		TotalPrice  FlexStr `json:"total_price"`
		NmID        FlexStr `json:"nm_id"`
		Brand       string  `json:"brand"`
		Status      FlexStr `json:"status"`
	} `json:"items"`
	Locale            string  `json:"locale"`
	InternalSignature string  `json:"internal_signature"`
	CustomerID        string  `json:"customer_id"`
	DeliveryService   string  `json:"delivery_service"`
	Shardkey          FlexStr `json:"shardkey"`
	SmID              FlexStr `json:"sm_id"`
	DateCreated       string  `json:"date_created"`
	OofShard          FlexStr `json:"oof_shard"`
}
