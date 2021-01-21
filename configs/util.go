package config

import (
	"os"
)

func readEnv(envName string) (string, error) {
	value := os.Getenv(envName)
	if value == "" {
		return "", ErrEnvNotFound{envName}
	}

	return value, nil
}
