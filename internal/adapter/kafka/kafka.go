package kafkaadapter

import (
	"fmt"

	"github.com/segmentio/kafka-go"

	"ModerationService/internal/config"
)

type (
	Topic string

	KafkaRepo struct {
		writers map[string]*kafka.Writer
		readers map[string]*kafka.Reader

		kafkaCfg *config.Kafka
	}
)

func (k *KafkaRepo) GetTopicByAlias(alias string) (Topic, error) {
	for _, t := range k.kafkaCfg.Topics {
		if t.Alias == alias {
			return Topic(t.Topic), nil
		}
	}
	return "", fmt.Errorf("alias not found: %s", alias)
}

func (t Topic) String() string {
	return string(t)
}
