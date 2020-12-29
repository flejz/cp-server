package cache

import (
	err "github.com/flejz/cp-server/internal/error"
)

type Memory struct {
	pair map[string]string
}

func (m *Memory) connect() error {
	m.pair = make(map[string]string)
	return nil
}

func (m *Memory) get(key string) (error, string) {
	value := m.pair[key]
	if value == "" {
		return &err.KeyNotFoundError{Key: key}, ""
	}

	return nil, m.pair[key]
}

func (m *Memory) set(key string, value string) error {
	m.pair[key] = value
	return nil
}
