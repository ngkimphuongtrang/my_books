package kafka

import (
	"errors"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/trangnkp/my_books/src/internal/kafka/shared"
)

type Producer struct {
	producer *kafka.Producer
}

type ProducerBuilder struct {
	shared.ParamBuilder
}

func NewProducerBuilder() *ProducerBuilder {
	return &ProducerBuilder{
		shared.ParamBuilder{
			Params: make(map[string]interface{}),
		},
	}
}

func (b *ProducerBuilder) WithBootstrapServers(brokers string) *ProducerBuilder {
	b.ParamBuilder.WithBootstrapServers(brokers)
	return b
}

func (b *ProducerBuilder) Build() (*Producer, error) {
	config := b.ToKafkaConfig()
	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, err
	}

	return &Producer{producer: producer}, nil
}

func (p *Producer) Produce(topic string, key []byte, value []byte) (*kafka.Message, error) {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
		Key:            key,
	}

	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err := p.producer.Produce(message, deliveryChan)
	if err != nil {
		return nil, err
	}

	r := <-deliveryChan
	m, ok := r.(*kafka.Message)
	if !ok {
		return nil, errors.New("type assertion failed: kafka.Message")
	}

	if m.TopicPartition.Error != nil {
		return nil, m.TopicPartition.Error
	}

	return m, nil
}

// Close closes the producer object
func (p *Producer) Close() {
	p.producer.Close()
}
