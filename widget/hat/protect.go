package hat

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const ProtectTag = "protect"

type Protector interface {
	Protect(strategies ...string) error
}

// Protect protects the fields of a struct based on the provided strategies.
// It works by setting the fields to their zero values if they are not included in the strategies.
// The strategies are defined in the struct tags with the key "protect".
// The strategies are comma-separated values.
// Example:
//
//	type User struct {
//		ID    string // This field will not be protected
//		Name  string `protect:"for_admin"`
//		Email string `protect:"for_admin,for_user"`
//		Age   int    `protect:"for_user"`
//	}
//
//	user := User{
//		ID:    "123",
//		Name:  "John",
//		Email: "something@mail.com",
//		Age:   30,
//	}
//
//	err := Protect(&user, "for_admin")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(user)
//
// In this example, the Protect function will set the Age field to its zero value because it is not included in the "for_admin" strategy.
// The Name & Email field will not be changed because it is included in the "for_admin" strategy.
// The ID field will not be changed because it does not have the "protect" tag.
func Protect(data any, strategies ...string) error {
	ref := reflect.ValueOf(data)

	if ref.Kind() != reflect.Ptr {
		return errors.New("data must be a pointer")
	}

	if ref.IsNil() {
		return errors.New("data must not be nil")
	}

	strategyMap := sliceToMap(strategies)
	ref = ref.Elem()

	if protector, ok := data.(Protector); ok {
		if err := protector.Protect(strategies...); err != nil {
			return fmt.Errorf("failed to protect: %w", err)
		}
	}

	if ref.Kind() == reflect.Slice {
		for i := 0; i < ref.Len(); i++ {
			if err := Protect(ref.Index(i).Addr().Interface(), strategies...); err != nil {
				return err
			}
		}

		return nil
	}

	if ref.Kind() == reflect.Struct {
		for i := 0; i < ref.NumField(); i++ {
			field := ref.Field(i)
			typ := ref.Type().Field(i)
			visible := typ.IsExported()

			if !visible {
				continue
			}

			if field.Kind() == reflect.Ptr && field.IsNil() {
				continue
			}

			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			if err := Protect(field.Addr().Interface(), strategies...); err != nil {
				return fmt.Errorf("failed to protect field %s: %w", typ.Name, err)
			}

			tag := typ.Tag.Get(ProtectTag)

			if tag == "" {
				continue
			}

			if !field.CanSet() {
				return errors.New("field is not settable")
			}

			if !mapContains(strategyMap, strings.Split(tag, ",")) {
				field.Set(reflect.Zero(field.Type())) // Set the field to its zero value
			}
		}
	}

	return nil
}

func sliceToMap(slice []string) map[string]struct{} {
	m := make(map[string]struct{}, len(slice))

	for _, s := range slice {
		m[s] = struct{}{}
	}

	return m
}

func mapContains(m map[string]struct{}, keys []string) bool {
	for _, key := range keys {
		if _, ok := m[key]; ok {
			return true
		}
	}

	return false
}
