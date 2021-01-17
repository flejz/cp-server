package configs

import (
	"github.com/flejz/cp-server/internal/errors"
	"os"
	"strconv"
)

type ServiceConfig struct {
	Port int
}

func (s *ServiceConfig) Load() error {
	port, portErr := strconv.Atoi(os.Getenv("PORT"))

	if portErr != nil {
		return &errors.ServiceConfigLoadError{Prop: "Port"}
	}

	s.Port = port

	return nil
}
