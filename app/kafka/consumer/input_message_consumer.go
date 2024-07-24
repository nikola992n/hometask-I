package consumer

import (
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"hometask/app/config"
	"hometask/app/kafka/producer"
	"time"
)

type InputMessageConsumer struct {
	consumer *Consumer
	producer *producer.Producer
}

func NewInputMessageConsumer(cfg config.Config, p *producer.Producer) (*InputMessageConsumer, error) {
	c, err := NewConsumer(cfg)
	if err != nil {
		return nil, errors.Join(errors.New("failed to create inputMessageConsumer"), err)
	}
	if err = c.SubscribeToTopics([]string{cfg.Kafka.InputTopic}); err != nil {
		return nil, err
	}
	return &InputMessageConsumer{
		consumer: c,
		producer: p,
	}, nil
}

func (imc *InputMessageConsumer) ConsumeMessages(
	processMessage func(msg *kafka.Message, producer *producer.Producer, outputTopic string) error,
) {
	for {
		timeout := time.Duration(imc.consumer.cfg.Kafka.ConsumerTimeout) * time.Second
		msg, err := imc.consumer.kafkaConsumer.ReadMessage(timeout)
		if err != nil {
			fmt.Println(err)
			// do error handling here
			continue
		}

		err = processMessage(msg, imc.producer, imc.consumer.cfg.Kafka.OutputTopic)
		if err != nil {
			fmt.Println(err)
			// choose what to do with unprocessed message
		}

		// commit when processing is successful
		imc.consumer.kafkaConsumer.CommitMessage(msg)
	}
}

func (imc *InputMessageConsumer) CloseConsumer() error {
	return imc.consumer.kafkaConsumer.Close()
}
