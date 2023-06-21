package parser

import (
	"l0/internal/models"
	"reflect"
)

func valueConverter(value reflect.Value) any {
	if intValue, ok := value.Interface().(int64); ok {
		return intValue
	} else if floatValue, ok := value.Interface().(float64); ok {
		return floatValue
	} else if strValue, ok := value.Interface().(string); ok {
		return strValue
	} else if flexFloatValue, ok := value.Interface().(models.FlexFloat); ok {
		return flexFloatValue
	} else if flexString, ok := value.Interface().(models.FlexStr); ok {
		return flexString
	}
	return value.String()
}

func OrderStructToSlice(order *models.Order) []interface{} {
	orderV := reflect.ValueOf(order)
	if orderV.Kind() == reflect.Ptr {
		orderV = orderV.Elem()
	}
	deliveryV := orderV.FieldByName("Delivery")
	paymentV := orderV.FieldByName("Payment")

	structLen := (orderV.NumField() - 1) + (deliveryV.NumField() - 1) + (paymentV.NumField() - 1)
	orderSlc := make([]interface{}, structLen)
	orderIndex := 0

	for i := 0; i < orderV.NumField(); i++ {
		if orderV.Type().Field(i).Name != "Items" &&
			orderV.Type().Field(i).Name != "Payment" &&
			orderV.Type().Field(i).Name != "Delivery" {
			orderSlc[orderIndex] = valueConverter(orderV.Field(i))
			orderIndex += 1
		}
	}
	for i := 0; i < deliveryV.NumField(); i++ {
		orderSlc[orderIndex] = valueConverter(deliveryV.Field(i))
		orderIndex += 1
	}
	for i := 0; i < paymentV.NumField(); i++ {
		orderSlc[orderIndex] = valueConverter(paymentV.Field(i))
		orderIndex += 1
	}

	return orderSlc
}

func ItemStructToSlice(i int, order *models.Order, userOrderId int64) []interface{} {
	itemV := reflect.ValueOf(order.Items[i])
	if itemV.Kind() == reflect.Ptr {
		itemV = itemV.Elem()
	}

	itemScl := make([]interface{}, itemV.NumField()+1)

	itemScl[0] = userOrderId
	for j := 0; j < itemV.NumField(); j++ {
		itemScl[j+1] = valueConverter(itemV.Field(j))
	}
	return itemScl
}
