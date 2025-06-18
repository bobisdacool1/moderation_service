package transport

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, h *Handlers) {
	app.Get("/health", h.healthcheckHandler.Healthcheck)

	api := app.Group("/api")

	api.Post("/moderation", h.moderationRequestHandler.CreateHandler)
	api.Post("/moderation/:id/approve", h.moderationRequestHandler.ApproveHandler)
	api.Post("/moderation/:id/decline", h.moderationRequestHandler.DeclineHandler)
}
