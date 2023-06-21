package parser

import (
	"fmt"
	"reflect"
)

func GetFieldIndex(structure interface{}, fieldName string) (int, error) {
	structValue := reflect.ValueOf(structure)
	if structValue.Kind() != reflect.Struct {
		return -1, fmt.Errorf("structure argument must be a struct")
	}

	structType := structValue.Type()
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Name == fieldName {
			return i, nil
		}
	}

	return -1, fmt.Errorf("field '%s' not found in the structure", fieldName)
}
