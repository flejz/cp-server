package db

import (
	"os"
)

type Type int

const defaultSQLitePath = "cp-server.db"
const (
	Memory Type = iota
	SQLite
)

type Config struct {
	Type       Type
	SQLitePath string
	MemoryName string
}

func Load(defaults bool) (*Config, error) {
	if defaults {
		os.Setenv("DB_TYPE", "sqlite")
		os.Setenv("SQLITE_PATH", defaultSQLitePath)
	}

	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		return nil, ErrInvalidType
	}

	var config *Config

	switch dbType {
	case "mem":
		config = &Config{Type: Memory, MemoryName: os.Getenv("MEM_NAME")}

		if config.MemoryName == "" {
			return nil, ErrInvalidMemoryName
		}
	case "sqlite":
		config = &Config{Type: SQLite, SQLitePath: os.Getenv("SQLITE_PATH")}

		if config.SQLitePath == "" {
			return nil, ErrInvalidSQLitePath
		}
	}

	return config, nil
}
