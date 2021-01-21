package model

import (
	"github.com/flejz/cp-server/internal/cache"
)

type Buffer struct {
	Cache cache.CacheInterface
}

func (self *Buffer) Get(usr, key string) (string, error) {
	return self.Cache.Get(usr, key)
}

func (self *Buffer) Set(usr, key, value string) error {
	return self.Cache.Set(usr, key, value)
}
