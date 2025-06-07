package hat

import (
	"errors"
	"reflect"
)

func MapUpdates(v interface{}, tag string) (map[string]interface{}, error) {
	reflected := reflect.ValueOf(v)

	if reflected.Kind() == reflect.Ptr {
		reflected = reflected.Elem()
	}

	if reflected.Kind() != reflect.Struct {
		return nil, errors.New("expected struct")
	}

	updates := make(map[string]interface{})

	for i := 0; i < reflected.NumField(); i++ {
		field := reflected.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		key := reflected.Type().Field(i).Tag.Get(tag)
		if key == "" {
			continue
		}

		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		// TODO: Omit non "present" fields

		updates[key] = field.Interface()
	}

	return updates, nil
}
