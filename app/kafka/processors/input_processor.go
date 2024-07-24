package processors

import (
	"encoding/json"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"hometask/app/kafka/producer"
	"hometask/app/models"
	"hometask/app/validators"
)

func InputMessageProcessor(msg *kafka.Message, producer *producer.Producer, outputTopic string) error {
	// unmarshall
	var inputMsgRaw models.InputMessageRaw
	if err := json.Unmarshal(msg.Value, &inputMsgRaw); err != nil {
		return errors.Join(errors.New("JSON unmarshall error"), err)
	}

	inputMsg, err := inputMsgRaw.ToInputMessage()
	if err != nil {
		return errors.Join(errors.New("message mapping error"), err)

	}

	// validate
	isValid := validators.ValidateInputMessage(*inputMsg)

	// marshall output message
	outputMsg := models.NewOutputMessage(inputMsgRaw, isValid)
	outputMsgBytes, err := json.Marshal(outputMsg)

	// produce to output topic
	if err = producer.ProduceToTopic(outputTopic, outputMsgBytes); err != nil {
		return errors.Join(errors.New("output message producing error"), err)
	}

	return nil
}
