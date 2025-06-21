package inmemadapter

import (
	"context"

	"go.uber.org/fx"

	"ModerationService/internal/config"
)

func ProvideInMemCache(
	lc fx.Lifecycle,
	cfg *config.Config,
) *Adapter {
	cache, err := NewInMemModerationRequestAdapter(cfg)
	if err != nil {
		panic(err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			cache.Start(ctx)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			cache.Stop()
			return nil
		},
	})

	return cache
}
