package service

import (
	"fmt"
	"messagio/internal/kafka"
	"messagio/internal/repository"
)

type MessageService struct {
	MessageRepository *repository.MessageRepository
	KafkaProducer     *kafka.KafkaProducer
	KafkaConsumer     *kafka.KafkaConsumer
}

func NewMessageService(mr *repository.MessageRepository, kp *kafka.KafkaProducer, kc *kafka.KafkaConsumer) *MessageService {
	return &MessageService{
		MessageRepository: mr,
		KafkaProducer:     kp,
		KafkaConsumer:     kc,
	}
}

func (ms *MessageService) AddMessage(text string) (int, error) {
	id, err := ms.MessageRepository.InsertMessage(text)
	if err != nil {
		return 0, err
	}

	err = ms.KafkaProducer.SendMessage(fmt.Sprintf("%d", id), text)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ms *MessageService) UpdateMessageStatus(id int, status string) error {
	return ms.MessageRepository.UpdateMessageStatus(id, status)
}

func (ms *MessageService) GetMessageStats() (map[string]int, error) {
	return ms.MessageRepository.GetMessageStats()
}
