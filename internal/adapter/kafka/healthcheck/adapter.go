package healthcheckadapter

import (
	"context"

	kafkaadapter "ModerationService/internal/adapter/kafka"
)

type (
	Adapter struct {
		kafkaRepo *kafkaadapter.KafkaClient
	}
)

func NewHealthCheckAdapter(kafkaRepo *kafkaadapter.KafkaClient) *Adapter {
	return &Adapter{
		kafkaRepo: kafkaRepo,
	}
}

func (a *Adapter) Ping(ctx context.Context) error {
	return a.kafkaRepo.Ping(ctx)
}
