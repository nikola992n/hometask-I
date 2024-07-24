package models

type OutputMessage struct {
	InputMessageRaw
	Status string `json:"status"`
}

func NewOutputMessage(inputMsg InputMessageRaw, isValid bool) OutputMessage {
	status := "invalid"
	if isValid {
		status = "valid"
	}

	return OutputMessage{
		InputMessageRaw: inputMsg,
		Status:          status,
	}
}
