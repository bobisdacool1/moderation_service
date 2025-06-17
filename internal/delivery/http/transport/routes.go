package transport

import (
	"github.com/gofiber/fiber/v2"

	"ModerationService/internal/delivery/http/handler"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/healthcheck", handler.Healthcheck)

	app.Get("/next", handler.NextHandler)
	app.Post("/approve/:id", handler.ApproveHandler)
	app.Post("/decline/:id", handler.DeclineHandler)
	app.Post("/create", handler.CreateHandler)
}
