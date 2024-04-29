package service

import (
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/kafka"
)

type KafkaNotificationService struct {
	*kafka.Producer
}

func NewKafkaNotificationService(cfg *config.KafkaProducerConfig) (*KafkaNotificationService, error) {
	producer, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	return &KafkaNotificationService{Producer: producer}, nil
}
