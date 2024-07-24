package service

import (
	"messagio/internal/config"
	"messagio/internal/kafka"
	"messagio/internal/repository"
)

type Service struct {
	MessageService *MessageService
}

func NewService(repo *repository.Repository, cfg *config.Config) *Service {

	kp := kafka.NewKafkaProducer(cfg.Kafka.Broker, cfg.Kafka.Topic)
	kc := kafka.NewKafkaConsumer(cfg.Kafka.Broker, cfg.Kafka.Topic, cfg.Kafka.GroupId, repo.MessageRepository)

	go kc.ReadMessages()

	return &Service{
		MessageService: NewMessageService(repo.MessageRepository, kp, kc),
	}
}
