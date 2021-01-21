package model

import (
	"github.com/flejz/cp-server/internal"
	"github.com/flejz/cp-server/internal/cache"
)

type Salt struct {
	Cache cache.CacheInterface
}

func (self *Salt) Generate(usr string) (string, error) {
	salt := util.Salt()

	if err := self.Cache.Set(usr, "", salt); err != nil {
		return "", err
	}

	return salt, nil
}

func (self *Salt) Get(usr string) (string, error) {
	return self.Cache.Get(usr, "")
}
