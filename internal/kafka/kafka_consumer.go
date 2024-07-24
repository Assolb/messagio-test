package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"log"
	"messagio/internal/repository"
	"strconv"
)

type KafkaConsumer struct {
	Reader *kafka.Reader
	Repo   *repository.MessageRepository
}

func NewKafkaConsumer(broker string, topic string, groupID string, repo *repository.MessageRepository) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
	})
	return &KafkaConsumer{Reader: reader, Repo: repo}
}

func (kc *KafkaConsumer) ReadMessages() {
	for {
		msg, err := kc.Reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Println("Error reading message: ", err)
			continue
		}

		messageID, err := strconv.Atoi(string(msg.Key))
		if err != nil {
			log.Printf("Error converting message ID: %v", err)
			continue
		}
		messageText := string(msg.Value)

		log.Printf("Received message: ID=%s, Text=%s", messageID, messageText)

		if err := kc.Repo.UpdateMessageStatus(messageID, "processed"); err != nil {
			log.Printf("Error updating message status: %v", err)
		}
		logrus.Printf("Received message: %s", string(msg.Value))
	}
}
