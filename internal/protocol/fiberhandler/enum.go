package fiberhandler

import (
	"feature/internal/value/domain"

	"github.com/gofiber/fiber/v2"
)

func enumServe(enums []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(enums)
	}
}

var (
	SortableFields    = enumServe(domain.SortableFields)
	QueryableFields   = enumServe(domain.QueryableFields)
	GroupableFields   = enumServe(domain.GroupableFields)
	WithableFields    = enumServe(domain.WithableFields)
	AccumulableFields = enumServe(domain.AccumulableFields)
)
