package moderation_request

import (
	"context"
	"encoding/json"
	"fmt"

	"ModerationService/internal/entity"
)

func (a *Adapter) WriteModerationRequest(ctx context.Context, request *entity.ModerationRequest) (entity.KafkaMessageEnvelope, error) {
	msg, err := a.writeEntity(ctx, a.moderationRequestsTopic, request, request.ID)
	if err != nil {
		return entity.KafkaMessageEnvelope{}, fmt.Errorf("failed to write entity: %w", err)
	}

	return msg, nil
}

func (a *Adapter) GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, entity.KafkaMessageEnvelope, error) {
	message, err := a.kafkaRepo.ReadMessage(ctx, a.moderationRequestsTopic)
	if err != nil {
		return nil, entity.KafkaMessageEnvelope{}, fmt.Errorf("failed to read moderation request: %w", err)
	}

	var request entity.ModerationRequest
	err = json.Unmarshal(message.Value, &request)
	if err != nil {
		return nil, entity.KafkaMessageEnvelope{}, fmt.Errorf("failed to unmarshal moderation request: %w", err)
	}

	return &request, entity.KafkaMessageToEnvelope(message), nil
}

func (a *Adapter) CommitModerationRequest(ctx context.Context, message entity.KafkaMessageEnvelope) error {
	err := a.kafkaRepo.CommitMessage(ctx, a.moderationRequestsTopic, entity.KafkaEnvelopeToMessage(message))
	if err != nil {
		return fmt.Errorf("failed to commit moderation request: %w", err)
	}

	return nil
}
