package kafka

import (
	"github.com/ngkimphuongtrang/runkit/kafka"
	log "github.com/sirupsen/logrus"
	"github.com/trangnkp/my_books/src/config"
)

type Producer struct {
	*kafka.Producer

	topic   string
	brokers string
}

func NewProducer(cfg *config.KafkaProducerConfig) (*Producer, error) {
	builder := kafka.NewProducerBuilder().WithBootstrapServers(cfg.Broker.Brokers)
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
