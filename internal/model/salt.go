package model

import (
	"github.com/flejz/cp-server/internal"
	"github.com/flejz/cp-server/internal/cache"
)

type Salt struct {
	Cache cache.CacheInterface
}

func (self *Salt) Generate(key string) (string, error) {
	salt := util.Salt()

	if err := self.Cache.Set(key, "", salt); err != nil {
		return "", err
	}

	return salt, nil
}

func (self *Salt) Get(key string) (string, error) {
	return self.Cache.Get(key, "")
}
