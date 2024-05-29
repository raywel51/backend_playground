package credential_helper

import (
	"errors"
	"fmt"
	"reflect"
)

func IsEmpty(obj interface{}) error {
	if obj == nil {
		return errors.New("object is nil")
	}

	var emptyFields []string

	valueOf := reflect.ValueOf(obj).Elem()
	typeOf := valueOf.Type()

	for i := 0; i < valueOf.NumField(); i++ {
		fieldValue := valueOf.Field(i)
		fieldName := typeOf.Field(i).Name
		skipCheck := typeOf.Field(i).Tag.Get("skip")

		if skipCheck == "true" {
			continue
		}

		if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
			emptyFields = append(emptyFields, fieldName)
		}

		if fieldValue.Kind() == reflect.Bool && !fieldValue.Bool() {
			emptyFields = append(emptyFields, fieldName)
		}

		if isNumericZero(fieldValue) {
			emptyFields = append(emptyFields, fieldName)
		}
	}

	if len(emptyFields) > 0 {
		return fmt.Errorf("empty fields: %v", emptyFields)
	}

	return nil
}

func isNumericZero(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	default:
		return false
	}
}
