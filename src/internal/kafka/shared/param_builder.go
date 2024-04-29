package shared

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/trangnkp/my_books/src/internal/container"
)

type ParamBuilder struct {
	Params container.Map
}

func (b *ParamBuilder) WithBootstrapServers(servers string) *ParamBuilder {
	b.Params = b.Params.Merge(container.Map{
		"bootstrap.servers": servers,
	})

	return b
}

func (b *ParamBuilder) ToKafkaConfig() *kafka.ConfigMap {
	configMap := kafka.ConfigMap{}
	for k, v := range b.Params {
		configMap[k] = v
	}
	return &configMap
}
