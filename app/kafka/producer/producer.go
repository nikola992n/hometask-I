package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"hometask/app/config"
)

type Producer struct {
	cfg           config.Config
	kafkaProducer *kafka.Producer
}

func NewProducer(cfg config.Config) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Server,
	})
	if err != nil {
		return nil, err
	}
	return &Producer{
		cfg:           cfg,
		kafkaProducer: p,
	}, err
}

func (p *Producer) ProduceToTopic(topic string, msg []byte) error {
	return p.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)
}

func (p *Producer) CloseProducer() {
	p.kafkaProducer.Close()
}
