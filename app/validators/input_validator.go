package validators

import (
	"hometask/app/models"
	"time"
)

func ValidateInputMessage(im models.InputMessage) bool {
	// validate timestamp is not older than 24 hours
	now := time.Now()
	diff := im.Timestamp.Sub(now).Hours()
	if diff < -24 {
		return false
	}

	// validate data length is greater than 10 characters
	if len(im.Data) <= 10 {
		return false
	}

	return true
}
