package cache

import (
	"fmt"
	"github.com/flejz/cp-server/internal/errors"
)

type MemoryCache struct {
	Key  string
	pair map[string]string
}

func (m *MemoryCache) key(key string) string {
	return fmt.Sprintf("%s-%s", m.Key, key)
}

func (m *MemoryCache) Init() error {
	m.pair = make(map[string]string)
	return nil
}

func (m *MemoryCache) Get(key string) (string, error) {
	if key == "" {
		return "", &errors.KeyNotSetError{}
	}

	value := m.pair[m.key(key)]
	if value == "" {
		return "", &errors.KeyNotFoundError{Key: key}
	}

	return value, nil
}

func (m *MemoryCache) Set(key string, value string) error {
	if key == "" {
		return &errors.KeyNotSetError{}
	}

	if value == "" {
		return &errors.ValueNotSetError{}
	}

	m.pair[m.key(key)] = value
	return nil
}
