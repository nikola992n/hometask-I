package env

import (
	"fmt"
	"github.com/joho/godotenv"
)

const EnvFileName = ".env"

func LoadDotEnv() error {
	err := godotenv.Load(EnvFileName)
	if err != nil {
		return fmt.Errorf("error while loading .env file: %s", err.Error())
	}
	return nil
}
