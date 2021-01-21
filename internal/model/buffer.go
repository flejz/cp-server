package model

import (
	"github.com/flejz/cp-server/internal/cache"
)

type Buffer struct {
	Cache cache.CacheInterface
}

func (self *Buffer) Get(key string) (string, error) {
	return self.Cache.Get(key, "default")
}

func (self *Buffer) Set(key, value string) error {
	return self.Cache.Set(key, "default", value)
}
