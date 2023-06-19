package parser

import (
	"l0/internal/models"
	"reflect"
)

func OrderStructToSlice(order *models.Order) []string {
	orderV := reflect.ValueOf(order)
	if orderV.Kind() == reflect.Ptr {
		orderV = orderV.Elem()
	}
	deliveryV := orderV.FieldByName("Delivery")
	paymentV := orderV.FieldByName("Payment")

	structLen := (orderV.NumField() - 1) + (deliveryV.NumField() - 1) + (paymentV.NumField() - 1)
	orderSlc := make([]string, structLen)
	orderIndex := 0

	for i := 0; i < orderV.NumField(); i++ {
		if orderV.Type().Field(i).Name != "Items" &&
			orderV.Type().Field(i).Name != "Payment" &&
			orderV.Type().Field(i).Name != "Delivery" {
			orderSlc[orderIndex] = orderV.Field(i).String()
			orderIndex += 1
		}
	}
	for i := 0; i < deliveryV.NumField(); i++ {
		orderSlc[orderIndex] = deliveryV.Field(i).String()
		orderIndex += 1
	}
	for i := 0; i < paymentV.NumField(); i++ {
		orderSlc[orderIndex] = paymentV.Field(i).String()
		orderIndex += 1
	}
	return orderSlc
}

func ItemStructToSlice(i int, order *models.Order) []string {
	itemV := reflect.ValueOf(order.Items[i])
	if itemV.Kind() == reflect.Ptr {
		itemV = itemV.Elem()
	}

	itemScl := make([]string, itemV.NumField()+1)

	itemScl[0] = order.OrderUID
	for j := 0; j < itemV.NumField(); j++ {
		itemScl[j+1] = itemV.Field(j).String()
	}
	return itemScl
}
