package kafka

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trangnkp/my_books/src/config"
)

func TestProducer(t *testing.T) {
	t.Parallel()

	cfg := config.NewKafkaProducerConfig()
	producer, err := NewProducer(cfg)
	require.NoError(t, err)
	t.Cleanup(producer.Close)

	t.Run("produce", func(t *testing.T) {
		message, err := producer.Produce(cfg.Topic, []byte("key"), []byte("value"))
		require.NoError(t, err)
		require.Equal(t, "key", string(message.Key))
		require.Equal(t, "value", string(message.Value))
	})
}
