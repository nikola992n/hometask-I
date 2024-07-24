package models

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type InputMessageRaw struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
}

func (itr *InputMessageRaw) ToInputMessage() (*InputMessage, error) {
	id, err := uuid.Parse(itr.ID)
	if err != nil {
		return nil, errors.Join(errors.New("failed to parse UUID from string"), err)
	}

	tm, err := time.Parse(time.DateTime, itr.Timestamp)
	if err != nil {
		return nil, errors.Join(errors.New("failed to parse Time from string"), err)
	}

	return &InputMessage{
		ID:        id,
		Timestamp: tm,
		Data:      itr.Data,
	}, nil
}

type InputMessage struct {
	ID        uuid.UUID
	Timestamp time.Time
	Data      string
}
