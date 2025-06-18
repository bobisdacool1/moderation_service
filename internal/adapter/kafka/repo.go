package kafkaadapter

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func (k *KafkaRepo) WriteMessage(ctx context.Context, topic string, key, value []byte) error {
	writer, ok := k.writers[topic]
	if !ok {
		return fmt.Errorf("no writer for topic: %s", topic)
	}
	return writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	})
}

func (k *KafkaRepo) ReadMessage(ctx context.Context, topic string) (kafka.Message, error) {
	reader, ok := k.readers[topic]
	if !ok {
		return kafka.Message{}, fmt.Errorf("no reader for topic: %s", topic)
	}
	return reader.FetchMessage(ctx)
}

func (k *KafkaRepo) CommitMessage(ctx context.Context, topic string, msg kafka.Message) error {
	reader, ok := k.readers[topic]
	if !ok {
		return fmt.Errorf("no reader for topic: %s", topic)
	}
	return reader.CommitMessages(ctx, msg)
}

func (k *KafkaRepo) Ping(ctx context.Context) error {
	conn, err := kafka.DialContext(ctx, "tcp", k.kafkaCfg.Broker)
	if err != nil {
		return fmt.Errorf("kafka ping failed: %w", err)
	}
	defer conn.Close()

	_, err = conn.ReadPartitions()
	if err != nil {
		return fmt.Errorf("kafka read partitions failed: %w", err)
	}

	return nil
}
