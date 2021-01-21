package tcp

import (
	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/errors"
	"github.com/flejz/cp-server/internal/user"
	"strings"
)

type Session struct {
	BufferModel buffer.BufferModel
	UserModel   user.UserModel
	logged      bool
	usr         string
}

func (self *Session) Register(usr, pwd string) error {
	return self.UserModel.Register(usr, pwd)
}

func (self *Session) Login(usr, pwd string) error {
	if err := self.UserModel.Validate(usr, pwd); err != nil {
		return err
	}

	self.logged = true
	self.usr = usr

	return nil
}

func (self *Session) Logout() error {
	self.logged = false
	self.usr = ""

	return nil
}

func (self *Session) Get(args []string) (string, error) {
	if !self.logged || len(args) > 1 {
		return "", &errors.InvalidError{}
	}

	key := ""

	if len(args) > 0 {
		key = args[0]
	}

	val, _ := self.BufferModel.Get(self.usr, key)
	return val, nil
}

func (self *Session) Set(args []string) error {
	if !self.logged || len(args) < 1 {
		return &errors.InvalidError{}
	}

	key := ""
	index := 0

	if len(args) > 1 {
		key = args[0]
		index = 1
	}
	return self.BufferModel.Set(self.usr, key, strings.Join(args[index:], " "))
}
