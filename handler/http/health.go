// Package http ...
package http

import "github.com/gofiber/fiber/v2"

// HealthCheck ...
func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
