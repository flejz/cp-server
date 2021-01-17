package cache

import (
	"database/sql"
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/errors"
	"github.com/mattn/go-sqlite3"
	"os"
)

type SQLite struct {
	Key    string
	config configs.ServerConfig
}

func (s *SQLite) Init() error {
	if _, err := os.Stat(s.Config.SQLitePath); os.IsNotExist(err) {
		file, err := os.Create(s.Config.SQLitePath)
		if err != nil {
			panic(err)
		}
		file.Close()
	}
}

func (s *SQLite) Get(key string) (string, error) {
	if key == "" {
		return "", &errors.KeyNotSetError{}
	}

	value := m.pair[m.key(key)]
	if value == "" {
		return "", &errors.KeyNotFoundError{Key: key}
	}

	return value, nil
}

func (s *SQLite) Set(key string, value string) error {
	if key == "" {
		return &errors.KeyNotSetError{}
	}

	if value == "" {
		return &errors.ValueNotSetError{}
	}

	m.pair[m.key(key)] = value
	return nil
}
