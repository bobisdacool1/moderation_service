package moderation_request

import (
	"context"
	"encoding/json"
	"fmt"

	"ModerationService/internal/entity"
)

type (
	ModerationRequestAdapter interface {
		WriteModerationRequest(ctx context.Context, request *entity.ModerationRequest) (entity.KafkaMessageEnvelope, error)
		GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, entity.KafkaMessageEnvelope, error)
		CommitModerationRequest(ctx context.Context, message entity.KafkaMessageEnvelope) error
		WriteApprovedRequest(ctx context.Context, request *entity.ModerationRequest) (entity.KafkaMessageEnvelope, error)
		WriteDeclinedRequest(ctx context.Context, request *entity.ModerationRequest) (entity.KafkaMessageEnvelope, error)
	}

	InMemAdapter interface {
		Get(ctx context.Context, id string) (entity.KafkaMessageEnvelope, bool)
		Put(ctx context.Context, id string, message entity.KafkaMessageEnvelope) error
		Delete(ctx context.Context, id string)
	}

	ModerationRequestService struct {
		moderationRequestAdapter ModerationRequestAdapter
		inMemRequestAdapter      InMemAdapter
	}
)

func NewModerationRequestService(moderationRequestAdapter ModerationRequestAdapter, inMemAdapter InMemAdapter) *ModerationRequestService {
	return &ModerationRequestService{
		moderationRequestAdapter: moderationRequestAdapter,
		inMemRequestAdapter:      inMemAdapter,
	}
}

func (s *ModerationRequestService) CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	_, err := s.moderationRequestAdapter.WriteModerationRequest(ctx, request)
	if err != nil {
		return fmt.Errorf("create request service: %w", err)
	}

	return nil
}

func (s *ModerationRequestService) GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, error) {
	moderationRequest, msg, err := s.moderationRequestAdapter.GetModerationRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("create request service: %w", err)
	}

	err = s.inMemRequestAdapter.Put(ctx, moderationRequest.ID, msg)
	if err != nil {
		return nil, fmt.Errorf("put request service: %w", err)
	}

	return moderationRequest, nil
}

func (s *ModerationRequestService) ReleaseModerationRequest(ctx context.Context, id string) (*entity.ModerationRequest, error) {
	msg, ok := s.inMemRequestAdapter.Get(ctx, id)
	if !ok {
		return nil, fmt.Errorf("failed to get stored moderation request")
	}

	err := s.moderationRequestAdapter.CommitModerationRequest(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("commit moderation request service: %w", err)
	}

	var request entity.ModerationRequest
	err = json.Unmarshal(msg.Value(), &request)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal moderation request: %w", err)
	}

	s.inMemRequestAdapter.Delete(ctx, id)

	return &request, nil
}

func (s *ModerationRequestService) CreateApprovedRequest(ctx context.Context, request *entity.ModerationRequest) error {
	_, err := s.moderationRequestAdapter.WriteApprovedRequest(ctx, request)
	if err != nil {
		return fmt.Errorf("create approved request service: %w", err)
	}

	return nil
}

func (s *ModerationRequestService) CreateDeclinedRequest(ctx context.Context, request *entity.ModerationRequest) error {
	_, err := s.moderationRequestAdapter.WriteDeclinedRequest(ctx, request)
	if err != nil {
		return fmt.Errorf("create approved request service: %w", err)
	}

	return nil
}
