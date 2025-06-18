package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	kafkaadapter "ModerationService/internal/adapter/kafka"
	moderationRequestAdapter "ModerationService/internal/adapter/kafka/moderation_request"
	"ModerationService/internal/config"
	"ModerationService/internal/delivery/http/handler"
	"ModerationService/internal/delivery/http/middleware"
	"ModerationService/internal/delivery/http/transport"
	moderationRequestSerivce "ModerationService/internal/service/moderation_request"
	moderationRequestUsecase "ModerationService/internal/usecases/moderation_request"
)

func NewFiberApp(cfg *config.Config) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	fiberApp.Use(middleware.RecoveryMiddleware())

	return fiberApp
}

func NewApp() *fx.App {
	providers := []interface{}{
		config.NewConfig,
		NewFiberApp,
		kafkaadapter.NewKafkaRepo,
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
	}

	invokes := []interface{}{
		func(lc fx.Lifecycle, k *kafkaadapter.KafkaRepo) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return k.EnsureTopics(ctx)
				},
			})
		},
		transport.StartHTTPServer,
	}

	return fx.New(
		fx.Provide(providers...),
		fx.Invoke(invokes...),
	)
}
