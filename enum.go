package enum

import (
	"reflect"
	"strconv"
)

func New[T any]() (enum T) {
	enumValue := reflect.ValueOf(&enum).Elem()
	enumType := reflect.TypeOf(&enum).Elem()
	var intCursor int64
	for i := 0; i < enumValue.NumField(); i++ {
		field := enumValue.Field(i)
		fieldType := enumType.Field(i)
		if !field.CanSet() {
			intCursor++
			continue
		}
		customValue := fieldType.Tag.Get("enum")
		switch k := field.Kind(); k {
		default:
			panic("enum field must be string or integer")
		case reflect.String:
			if customValue != "" {
				field.SetString(customValue)
			} else {
				field.SetString(fieldType.Name)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			customInt, err := strconv.ParseInt(customValue, 10, 64)
			if err == nil {
				intCursor = customInt
			}
			field.SetInt(intCursor)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			customUint, err := strconv.ParseUint(customValue, 10, 64)
			if err == nil {
				intCursor = int64(customUint)
			}
			field.SetUint(uint64(intCursor))
		}
		intCursor++
	}
	return enum
}
