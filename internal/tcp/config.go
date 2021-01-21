package tcp

import (
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func Load() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return nil, ErrInvalidPort
	}

	return &Config{port}, nil
}
