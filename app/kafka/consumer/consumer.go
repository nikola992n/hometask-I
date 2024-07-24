package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"hometask/app/config"
)

type Consumer struct {
	cfg           config.Config
	kafkaConsumer *kafka.Consumer
}

func NewConsumer(cfg config.Config) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  cfg.Kafka.Server,
		"group.id":           cfg.Kafka.GroupID,
		"auto.offset.reset":  kafka.OffsetBeginning,
		"enable.auto.commit": false,
	})
	if err != nil {
		return nil, err
	}

	return &Consumer{
		cfg:           cfg,
		kafkaConsumer: c,
	}, nil
}

func (c *Consumer) SubscribeToTopics(topics []string) error {
	return c.kafkaConsumer.SubscribeTopics(topics, nil)
}
