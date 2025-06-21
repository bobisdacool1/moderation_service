package kafkaadapter

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func (k *KafkaClient) WriteMessage(ctx context.Context, topic Topic, key, value []byte) (kafka.Message, error) {
	writer, ok := k.writers[topic.String()]
	if !ok {
		return kafka.Message{}, fmt.Errorf("no writer for topic: %s", topic)
	}

	msg := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	err := writer.WriteMessages(ctx, msg)
	if err != nil {
		return kafka.Message{}, fmt.Errorf("failed to write message: %w", err)
	}

	return msg, nil
}

func (k *KafkaClient) ReadMessage(ctx context.Context, topic Topic) (kafka.Message, error) {
	reader, ok := k.readers[topic.String()]
	if !ok {
		return kafka.Message{}, fmt.Errorf("no reader for topic: %s", topic)
	}
	return reader.FetchMessage(ctx)
}

func (k *KafkaClient) CommitMessage(ctx context.Context, topic Topic, msg kafka.Message) error {
	reader, ok := k.readers[topic.String()]
	if !ok {
		return fmt.Errorf("no reader for topic: %s", topic)
	}
	return reader.CommitMessages(ctx, msg)
}

func (k *KafkaClient) Ping(ctx context.Context) error {
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
