package cache

import (
	"fmt"
	err "github.com/flejz/cp-server/internal/error"
)

type Memory struct {
	Key  string
	pair map[string]string
}

func (m *Memory) key(key string) string {
	return fmt.Sprintf("%s-%s", m.Key, key)
}

func (m *Memory) Init() error {
	m.pair = make(map[string]string)
	return nil
}

func (m *Memory) Get(key string) (string, error) {
	if key == "" {
		return "", &err.KeyNotSetError{}
	}

	value := m.pair[m.key(key)]
	if value == "" {
		return "", &err.KeyNotFoundError{Key: key}
	}

	return value, nil
}

func (m *Memory) Set(key string, value string) error {
	if key == "" {
		return &err.KeyNotSetError{}
	}

	if value == "" {
		return &err.ValueNotSetError{}
	}

	m.pair[m.key(key)] = value
	return nil
}
