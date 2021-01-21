package cache

import (
	"fmt"
	"github.com/flejz/cp-server/internal/errors"
)

type MemoryCache struct {
	BaseCache
	Key  string
	pair map[string]string
}

func (m *MemoryCache) usrKey(usr, key string) string {
	return fmt.Sprintf("%s-%s-%s", m.Key, usr, m.key(key))
}

func (m *MemoryCache) Init() error {
	m.pair = make(map[string]string)
	return nil
}

func (m *MemoryCache) Get(usr, key string) (string, error) {
	if usr == "" {
		return "", &errors.KeyNotSetError{}
	}

	value := m.pair[m.usrKey(usr, key)]
	if value == "" {
		return "", &errors.KeyNotFoundError{Key: key}
	}

	return value, nil
}

func (m *MemoryCache) Set(usr, key, value string) error {
	if usr == "" {
		return &errors.KeyNotSetError{}
	}

	if value == "" {
		return &errors.ValueNotSetError{}
	}

	m.pair[m.usrKey(usr, key)] = value
	return nil
}
