package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	inmemadapter "ModerationService/internal/adapter/in_mem/moderation_request"
	kafkaadapter "ModerationService/internal/adapter/kafka"
	"ModerationService/internal/adapter/kafka/healthcheck"
	moderationRequestAdapter "ModerationService/internal/adapter/kafka/moderation_request"
	"ModerationService/internal/config"
	"ModerationService/internal/delivery/http/handler"
	"ModerationService/internal/delivery/http/middleware"
	"ModerationService/internal/delivery/http/transport"
	healthcheckservice "ModerationService/internal/service/healthcheck"
	moderationRequestSerivce "ModerationService/internal/service/moderation_request"
	healthcheckusecase "ModerationService/internal/usecases/healthcheck"
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
		kafkaadapter.NewKafkaClient,

		inmemadapter.ProvideInMemCache,
		func(c *inmemadapter.Adapter) moderationRequestSerivce.InMemAdapter {
			return c
		},
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

		fx.Annotate(
			healthcheckadapter.NewHealthCheckAdapter,
			fx.As(new(healthcheckservice.HealthcheckAdapter)),
		),
		fx.Annotate(
			healthcheckservice.NewHealthcheckService,
			fx.As(new(healthcheckusecase.HealthcheckService)),
		),
		fx.Annotate(
			healthcheckusecase.NewHealthcheckUsecase,
			fx.As(new(handler.HealthcheckUsecase)),
		),

		handler.NewModerationRequestHandler,
		handler.NewHealthcheckHandler,
		transport.NewHandlers,
	}

	invokes := []interface{}{
		func(lc fx.Lifecycle, k *kafkaadapter.KafkaClient) {
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
