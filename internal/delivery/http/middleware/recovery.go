package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
)

func RecoveryMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic recovered", zap.Any("panic", r), zap.Stack("stack"))
				_ = c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		}()

		return c.Next()
	}
}
