package moderation_request

import (
	"context"
	"encoding/json"
	"fmt"

	kafkaadapter "ModerationService/internal/adapter/kafka"
	"ModerationService/internal/entity"
)

type (
	Adapter struct {
		kafkaRepo               *kafkaadapter.KafkaRepo
		moderationRequestsTopic kafkaadapter.Topic
	}
)

func NewModerationRequestAdapter(kafkaRepo *kafkaadapter.KafkaRepo) *Adapter {
	moderationRequestTopic, err := kafkaRepo.GetTopicByAlias("moderation-requests")
	if err != nil {
		panic(err)
	}

	return &Adapter{
		kafkaRepo:               kafkaRepo,
		moderationRequestsTopic: moderationRequestTopic,
	}
}

func (a *Adapter) WriteModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	bb, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal moderation request: %w", err)
	}

	err = a.kafkaRepo.WriteMessage(ctx, a.moderationRequestsTopic.String(), a.getKey(request), bb)
	if err != nil {
		return fmt.Errorf("failed to write moderation request: %w", err)
	}

	return nil
}

func (a *Adapter) GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, error) {
	e, err := a.kafkaRepo.ReadMessage(ctx, a.moderationRequestsTopic.String())
	if err != nil {
		return nil, fmt.Errorf("failed to read moderation request: %w", err)
	}

	var request entity.ModerationRequest
	err = json.Unmarshal(e.Value, &request)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal moderation request: %w", err)
	}

	return &request, nil
}

func (a *Adapter) getKey(request *entity.ModerationRequest) []byte {
	return []byte(a.getEntityPrefix() + request.ID)
}

func (a *Adapter) getEntityPrefix() string {
	return a.moderationRequestsTopic.String() + ":"
}
