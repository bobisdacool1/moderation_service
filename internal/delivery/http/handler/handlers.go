package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeclineHandler(c *fiber.Ctx) error {
	return c.SendString("declined")
}

func NextHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"id": "abc123", "text": "пример"})
}

func ApproveHandler(c *fiber.Ctx) error {
	return c.SendString("approved")
}

func CreateHandler(c *fiber.Ctx) error {
	return c.SendString("created")
}

func Healthcheck(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
