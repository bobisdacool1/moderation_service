package transport

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"ModerationService/internal/delivery/http/handler"
)

type Handlers struct {
	healthcheckHandler       *handler.HealthcheckHandler
	moderationRequestHandler *handler.ModerationRequestHandler
}

func NewHandlers(
	healthcheckHandler *handler.HealthcheckHandler,
	moderationRequestHandler *handler.ModerationRequestHandler,
) *Handlers {
	return &Handlers{
		healthcheckHandler:       healthcheckHandler,
		moderationRequestHandler: moderationRequestHandler,
	}
}

func StartHTTPServer(lc fx.Lifecycle, app *fiber.App, h *Handlers) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				RegisterRoutes(app, h)

				log.Println("HTTP server running on :3000")
				if err := app.Listen(":3000"); err != nil {
					log.Fatalf("fiber error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
