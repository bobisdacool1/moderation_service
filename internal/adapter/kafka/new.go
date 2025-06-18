package kafkaadapter

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"ModerationService/internal/config"
)

func NewKafkaRepo(cfg *config.Config) (*KafkaRepo, error) {
	if len(cfg.Kafka.Broker) == 0 {
		return nil, fmt.Errorf("no kafka broker provided")
	}

	k := &KafkaRepo{
		writers: make(map[string]*kafka.Writer),
		readers: make(map[string]*kafka.Reader),

		kafkaCfg: &cfg.Kafka,
	}

	for _, c := range cfg.Kafka.Topics {
		topic := c.Topic
		broker := cfg.Kafka.Broker

		k.writers[topic] = &kafka.Writer{
			Addr:         kafka.TCP(broker),
			Topic:        topic,
			Balancer:     &kafka.Hash{},
			RequiredAcks: kafka.RequireAll,
		}

		k.readers[topic] = kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{broker},
			GroupID:        c.GroupID,
			Topic:          topic,
			MinBytes:       1,
			MaxBytes:       10e6,
			CommitInterval: 0,
		})
	}

	return k, nil
}

func (k *KafkaRepo) Close() error {
	var err error
	for _, r := range k.readers {
		err = r.Close()
	}
	for _, w := range k.writers {
		err = w.Close()
	}
	return err
}

func (k *KafkaRepo) EnsureTopics(ctx context.Context) error {
	conn, err := kafka.DialContext(ctx, "tcp", k.kafkaCfg.Broker)
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka broker: %w", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("failed to get Kafka controller: %w", err)
	}

	ctrlAddr := fmt.Sprintf("%s:%d", controller.Host, controller.Port)
	ctrlConn, err := kafka.Dial("tcp", ctrlAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka controller: %w", err)
	}
	defer ctrlConn.Close()

	for _, t := range k.kafkaCfg.Topics {
		err := ctrlConn.CreateTopics(kafka.TopicConfig{
			Topic:             t.Topic,
			NumPartitions:     t.NumPartitions,
			ReplicationFactor: t.ReplicationFactor,
		})
		if err != nil {
			return fmt.Errorf("failed to create topic %q: %w", t.Topic, err)
		}
	}

	return nil
}
