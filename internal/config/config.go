package config

import (
	"log"
	"os"
	"strconv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if s, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("bad value %q for %s: %v", s, key, err)
		}
		return v
	}
	return fallback
}

const (
	// MemoryMode is the memory db type
	MemoryMode = "mem"

	// SQLiteMode is the sqlite db type
	SQLiteMode = "sqlite"

	// internal
	memoryKeyDefault    = "cp-server"
	sQLiteDBPathDefault = "./cp-server.db"
)

// Config hold shared configuration values used for initating service components
type Config struct {
	DB  DBSettings
	TCP TCPSettings
}

// DBSettings hold db configuration
type DBSettings struct {
	Mode string
	Name string
	Path string
}

// TCPSettings holds the tcp configuration
type TCPSettings struct {
	Port int
}

// Init gets the configs for the application
func Init() (*Config, error) {
	return &Config{
		DB: DBSettings{
			Mode: getEnv("DB_MODE", SQLiteMode),
			Name: getEnv("DB_MEMORY_KEY", memoryKeyDefault),
			Path: getEnv("DB_SQLITE_PATH", sQLiteDBPathDefault),
		},
		TCP: TCPSettings{
			Port: getEnvInt("TCP_PORT", 2000),
		},
	}, nil
}
