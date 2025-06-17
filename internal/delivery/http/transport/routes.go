package transport

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, h *Handlers) {
	api := app.Group("/api")

	api.Get("/health", h.healthcheckHandler.Healthcheck)

	api.Post("/moderation", h.moderationRequestHandler.CreateHandler)
	api.Get("/moderation/next", h.moderationRequestHandler.NextHandler)
	api.Post("/moderation/:id/approve", h.moderationRequestHandler.ApproveHandler)
	api.Post("/moderation/:id/decline", h.moderationRequestHandler.DeclineHandler)
}
