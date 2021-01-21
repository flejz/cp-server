package configs

import (
	"os"
	"strconv"
)

type DBType int

const (
	Memory DBType = iota
	SQLite
)

type DBConfig struct {
	Type       DBType
	SQLitePath string
	MemoryName string
}

func (c *DBConfig) Load() error {
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		return ErrInvalidDBType
	}

	switch dbType {
	case "memory":
		c.Type = Memory
		c.MemoryName = os.Getenv("MEMORY_NAME")

	case "sqlite":
		c.Type = SQLite
	default:
		return error.New()
	}

	return nil
}
