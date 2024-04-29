package kafka

import (
	log "github.com/sirupsen/logrus"
	"github.com/trangnkp/my_books/src/config"
	sharedkafka "github.com/trangnkp/my_books/src/internal/kafka"
)

type Producer struct {
	*sharedkafka.Producer

	topic   string
	brokers string
}

func NewProducer(cfg *config.KafkaProducerConfig) (*Producer, error) {
	builder := sharedkafka.NewProducerBuilder().WithBootstrapServers(cfg.Broker.Brokers)
	producer, err := builder.Build()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Producer{
		topic:    cfg.Topic,
		brokers:  cfg.Broker.Brokers,
		Producer: producer,
	}, nil
}

func (p *Producer) Topic() string {
	return p.topic
}
