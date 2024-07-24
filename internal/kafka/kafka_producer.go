package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(broker string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: topic,
	}
	return &KafkaProducer{Writer: writer}
}

func (kp *KafkaProducer) SendMessage(key, value string) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}
	return kp.Writer.WriteMessages(context.Background(), msg)
}
