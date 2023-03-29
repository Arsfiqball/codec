package fiberhandler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
