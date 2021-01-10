package configs

import (
	"errors"
	"os"
	"strconv"
)

type DaemonConfig struct {
	Salt string
	Port int
}

func (d *DaemonConfig) Load() error {
	salt := os.Getenv("SALT")
	port, portErr := strconv.Atoi(os.Getenv("PORT"))

	if d.Salt == "" {
		return errors.New("Salt not defined")
	}

	if portErr != nil {
		return portErr
	}

	d.Salt = salt
	d.Port = port

	return nil
}
