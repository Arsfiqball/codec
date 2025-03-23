package hat

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func UnmarshalJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v) // Currently, we don't need to implement custom unmarshal logic
}

func MarshalJSON(v interface{}) ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	reflected := reflect.ValueOf(v)

	if reflected.Kind() == reflect.Ptr {
		reflected = reflected.Elem()
	}

	_, ok := v.(json.Marshaler)

	// If struct, iterate over fields
	if reflected.Kind() == reflect.Struct && !ok {
		str := "{"

		for i := 0; i < reflected.NumField(); i++ {
			b, err := MarshalJSON(reflected.Field(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("failed to marshal field %d: %w", i, err)
			}

			if string(b) == "{}" {
				continue // Skip empty fields
			}

			key := reflected.Type().Field(i).Tag.Get("json")
			if key == "" {
				key = reflected.Type().Field(i).Name
			}

			if i > 0 {
				str += ","
			}

			str += "\"" + key + "\":" + string(b)
		}

		str += "}"

		return []byte(str), nil
	}

	// Otherwise, marshal the value using standard json.Marshal
	return json.Marshal(v)
}
