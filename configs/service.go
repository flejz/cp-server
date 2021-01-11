package configs

import (
	err "github.com/flejz/cp-server/internal/error"
	"os"
	"strconv"
)

type ServiceConfig struct {
	Salt string
	Port int
}

func LoadServiceConfig() (*ServiceConfig, error) {
	salt := os.Getenv("SALT")
	port, portErr := strconv.Atoi(os.Getenv("PORT"))

	if salt == "" {
		return nil, &err.ServiceConfigLoadError{Prop: "Salt"}
	}

	if portErr != nil {
		return nil, &err.ServiceConfigLoadError{Prop: "Port"}
	}

	return &ServiceConfig{Salt: salt, Port: port}, nil
}
