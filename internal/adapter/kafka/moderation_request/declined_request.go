package moderation_request

import (
	"context"
	"fmt"

	"ModerationService/internal/entity"
)

func (a *Adapter) WriteDeclinedRequest(ctx context.Context, request *entity.ModerationRequest) (entity.KafkaMessageEnvelope, error) {
	msg, err := a.writeEntity(ctx, a.declinedRequestsTopic, request, request.ID)
	if err != nil {
		return entity.KafkaMessageEnvelope{}, fmt.Errorf("writeDeclinedRequest: %w", err)
	}

	return msg, nil
}
