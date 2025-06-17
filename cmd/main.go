package main

import (
	"go.uber.org/fx"

	"ModerationService/internal/app"
	"ModerationService/internal/delivery/http/transport"
)

func main() {
	fx.New(
		fx.Provide(
			app.NewFiberApp,
		),
		fx.Invoke(
			transport.StartHttpServer,
		),
	).Run()
}
