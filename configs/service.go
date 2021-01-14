package configs

import (
	"github.com/flejz/cp-server/internal/errors"
	"os"
	"strconv"
)

type ServiceConfig struct {
	Port int
}

func Load() (*ServiceConfig, error) {
	port, portErr := strconv.Atoi(os.Getenv("PORT"))

	if portErr != nil {
		return nil, &errors.ServiceConfigLoadError{Prop: "Port"}
	}

	return &ServiceConfig{Port: port}, nil
}
