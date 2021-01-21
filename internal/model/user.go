package model

import (
	"github.com/flejz/cp-server/internal"
	"github.com/flejz/cp-server/internal/cache"
	"github.com/flejz/cp-server/internal/errors"
)

type User struct {
	Cache     cache.CacheInterface
	SaltModel Salt
}

func (self *User) Register(usr, pwd string) error {
	salt, err := self.SaltModel.Generate(usr)
	if err != nil {
		return err
	}

	hash := util.Hash(pwd, salt)
	if err = self.Cache.Set(usr, "", hash); err != nil {
		return err
	}

	return nil
}

func (self *User) Validate(usr, pwd string) error {
	salt, err := self.SaltModel.Get(usr)
	if err != nil {
		return err
	}

	hash, err := self.Cache.Get(usr, "")
	if err != nil {
		return err
	} else if util.Hash(pwd, salt) != hash {
		return &errors.InvalidError{}
	}

	return nil
}
