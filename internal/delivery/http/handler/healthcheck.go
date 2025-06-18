package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	HealthcheckHandler struct{}
)

func NewHealthcheck() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Healthcheck(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
