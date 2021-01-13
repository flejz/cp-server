package configs

import (
	err "github.com/flejz/cp-server/internal/error"
	"os"
	"strconv"
)

type ServiceConfig struct {
	Port int
}

func Load() (*ServiceConfig, error) {
	port, portErr := strconv.Atoi(os.Getenv("PORT"))

	if portErr != nil {
		return nil, &err.ServiceConfigLoadError{Prop: "Port"}
	}

	return &ServiceConfig{Port: port}, nil
}
