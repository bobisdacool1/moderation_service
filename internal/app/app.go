package app

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	moderationRequestAdapter "ModerationService/internal/adapter/kafka/moderation_request"
	"ModerationService/internal/delivery/http/handler"
	"ModerationService/internal/delivery/http/transport"
	moderationRequestSerivce "ModerationService/internal/service/moderation_request"
	moderationRequestUsecase "ModerationService/internal/usecases/moderation_request"
)

func NewFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		AppName: "ModerationService",
	})
}

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			NewFiberApp,
			fx.Annotate(
				moderationRequestAdapter.NewModerationRequestAdapter,
				fx.As(new(moderationRequestSerivce.ModerationRequestAdapter)),
			),
			fx.Annotate(
				moderationRequestSerivce.NewModerationRequestService,
				fx.As(new(moderationRequestUsecase.ModerationRequestService)),
			),
			fx.Annotate(
				moderationRequestUsecase.NewModerationRequestUsecase,
				fx.As(new(handler.ModerationRequestUsecase)),
			),

			handler.NewModerationRequestHandler,

			handler.NewHealthcheck,

			transport.NewHandlers,
		),
		fx.Invoke(
			transport.StartHTTPServer,
		),
	)
}
