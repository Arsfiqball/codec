package fiberhandler

import "github.com/gofiber/fiber/v2"

func (h *Handler) BulkOps(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
