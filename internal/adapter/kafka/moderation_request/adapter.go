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
		kafkaRepo *kafkaadapter.KafkaClient

		moderationRequestsTopic kafkaadapter.Topic
		approvedRequestsTopic   kafkaadapter.Topic
		declinedRequestsTopic   kafkaadapter.Topic
	}
)

func NewModerationRequestAdapter(kafkaRepo *kafkaadapter.KafkaClient) *Adapter {
	moderationRequestTopic, err := kafkaRepo.GetTopicByAlias("moderation-requests")
	if err != nil {
		panic(err)
	}
	approvedRequestTopic, err := kafkaRepo.GetTopicByAlias("moderation-requests-approved")
	if err != nil {
		panic(err)
	}
	declinedRequestTopic, err := kafkaRepo.GetTopicByAlias("moderation-requests-declined")
	if err != nil {
		panic(err)
	}

	return &Adapter{
		kafkaRepo:               kafkaRepo,
		moderationRequestsTopic: moderationRequestTopic,
		approvedRequestsTopic:   approvedRequestTopic,
		declinedRequestsTopic:   declinedRequestTopic,
	}
}

func (a *Adapter) writeEntity(ctx context.Context, topic kafkaadapter.Topic, data any, id string) (entity.KafkaMessageEnvelope, error) {
	bb, err := json.Marshal(data)
	if err != nil {
		return entity.KafkaMessageEnvelope{}, fmt.Errorf("failed to marshal moderation request: %w", err)
	}

	msg, err := a.kafkaRepo.WriteMessage(ctx, topic, getKey(topic, id), bb)
	if err != nil {
		return entity.KafkaMessageEnvelope{}, fmt.Errorf("failed to write moderation request: %w", err)
	}

	return entity.KafkaMessageToEnvelope(msg), nil
}

func getKey(topic kafkaadapter.Topic, id string) []byte {
	return []byte(getEntityPrefix(topic) + id)
}

func getEntityPrefix(topic kafkaadapter.Topic) string {
	return topic.String() + ":"
}
