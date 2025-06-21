package moderation_request

import (
	"context"
	"fmt"

	"ModerationService/internal/entity"
)

func (a *Adapter) WriteApprovedRequest(ctx context.Context, request *entity.ModerationRequest) (entity.KafkaMessageEnvelope, error) {
	msg, err := a.writeEntity(ctx, a.approvedRequestsTopic, request, request.ID)
	if err != nil {
		return entity.KafkaMessageEnvelope{}, fmt.Errorf("writeApprovedRequest: %w", err)
	}

	return msg, nil
}
