package cfg

import (
	"reflect"

	"ttgoer/log"
)

func validateRequiredFields(cfg any) {
	v := reflect.ValueOf(cfg)
	t := reflect.TypeOf(cfg)

	if v.Kind() != reflect.Struct {
		log.S().Panic("Expected a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Struct {
			validateRequiredFields(field.Interface())
			continue
		}

		requiredTag := fieldType.Tag.Get("required")
		if requiredTag == "true" {
			if isZeroValue(field) {
				log.S().Panicf("field '%s' is required but has a zero value", fieldType.Name)
			}
		}
	}
}

func isZeroValue(field reflect.Value) bool {
	return reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface())
}
