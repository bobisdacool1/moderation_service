package transport

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func StartHttpServer(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				RegisterRoutes(app)

				log.Println("Server is running on :3000")
				if err := app.Listen(":3000"); err != nil {
					log.Printf("Fiber error: %v", err)
					os.Exit(1)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			return app.Shutdown()
		},
	})
}
