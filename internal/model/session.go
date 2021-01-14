package model

import (
	"github.com/flejz/cp-server/internal/errors"
	"strings"
)

type Session struct {
	BufferModel Buffer
	UserModel   User
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

func (self *Session) Get() (string, error) {
	if !self.logged {
		return "", &errors.InvalidError{}
	}

	val, _ := self.BufferModel.Get(self.usr)
	return val, nil
}

func (self *Session) Set(vals []string) error {
	if !self.logged || len(vals) < 1 {
		return &errors.InvalidError{}
	}

	return self.BufferModel.Set(self.usr, strings.Join(vals, " "))
}
