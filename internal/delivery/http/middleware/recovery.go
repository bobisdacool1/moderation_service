package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RecoveryMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic recovered")
				_ = c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		}()

		return c.Next()
	}
}
