package fiberhandler

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetOne(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
