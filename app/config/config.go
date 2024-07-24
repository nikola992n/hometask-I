package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"hometask/app/env"
)

type Kafka struct {
	Server          string `env:"KAFKA_SERVER" env-default:"localhost:9093"`
	GroupID         string `env:"KAFKA_GROUP_ID" env-default:"default-group"`
	ConsumerTimeout int    `env:"KAFKA_CONSUME_TIMEOUT" env-default:"10"`
	InputTopic      string `env:"KAFKA_TOPIC_INPUT" env-default:"default-topic"`
	OutputTopic     string `env:"KAFKA_TOPIC_OUTPUT" env-default:"default-topic"`
}

type Config struct {
	Kafka Kafka
}

func Load() (*Config, error) {
	if err := env.LoadDotEnv(); err != nil {
		return nil, err
	}

	var kafka Kafka
	if err := cleanenv.ReadEnv(&kafka); err != nil {
		return nil, err
	}

	return &Config{
		Kafka: kafka,
	}, nil
}
