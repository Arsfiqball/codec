package fiberhandler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Update(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
