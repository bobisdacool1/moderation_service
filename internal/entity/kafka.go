package entity

import (
	"github.com/segmentio/kafka-go"
)

type KafkaMessageEnvelope struct {
	raw kafka.Message
}

func KafkaEnvelopeToMessage(env KafkaMessageEnvelope) kafka.Message {
	return env.raw
}

func KafkaMessageToEnvelope(msg kafka.Message) KafkaMessageEnvelope {
	return KafkaMessageEnvelope{
		raw: msg,
	}
}

func (e KafkaMessageEnvelope) Value() []byte {
	return e.raw.Value
}
