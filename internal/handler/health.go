package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Health() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.JSON(&fiber.Map{"status": "HEALTHY"})
		return nil
	}
}
