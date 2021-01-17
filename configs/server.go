package configs

import (
	"github.com/flejz/cp-server/internal/errors"
	"os"
	"strconv"
)

type ServiceConfig struct {
	SQLitePath string
	Port       int
}

func (c *ServiceConfig) Load() error {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return &errors.ServiceConfigLoadError{Prop: "Port"}
	}

	sqlitePath := os.Getenv("SQLITE_PATH")

	if sqlitePath == "" {
		return &errors.ServiceConfigLoadError{Prop: "SQLite Path"}
	}

	c.Port = port
	c.SQLitePath = sqlitePath

	return nil
}
