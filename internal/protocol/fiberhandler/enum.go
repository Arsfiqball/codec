package fiberhandler

import (
	"github.com/Arsfiqball/codec/internal/value/domain"

	"github.com/gofiber/fiber/v2"
)

var fieldLabel = map[string]string{
	domain.FieldID:       "ID",
	domain.FieldName:     "Name",
	domain.FieldEmail:    "Email",
	domain.FieldPassword: "Password",
	domain.FieldCount:    "Count",
}

func labelizeEnumFields(fields []string) map[string]string {
	enums := make(map[string]string)

	for _, field := range fields {
		enums[field] = fieldLabel[field]
	}

	return enums
}

func enumServe(enums map[string]string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return FormatSuccess(c, enums)
	}
}

var (
	SortableFields    = enumServe(labelizeEnumFields(domain.SortableFields))
	QueryableFields   = enumServe(labelizeEnumFields(domain.QueryableFields))
	GroupableFields   = enumServe(labelizeEnumFields(domain.GroupableFields))
	WithableFields    = enumServe(labelizeEnumFields(domain.WithableFields))
	AccumulableFields = enumServe(labelizeEnumFields(domain.AccumulableFields))
)
