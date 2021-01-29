package buffer

import (
	"strings"

	"github.com/flejz/cp-server/internal/user"
)

type Cmd struct {
	buffMdl BufferModel
	usrMdl  user.UserModel
	logged  bool
	usr     string
}

func (self *Cmd) Register(usr, pwd string) error {
	return self.usrMdl.Register(usr, pwd)
}

func (self *Cmd) Login(usr, pwd string) error {
	if err := self.usrMdl.Validate(usr, pwd); err != nil {
		return err
	}

	self.logged = true
	self.usr = usr

	return nil
}

func (self *Cmd) Logout() error {
	self.logged = false
	self.usr = ""

	return nil
}

func (self *Cmd) Get(args []string) (string, error) {
	if !self.logged || len(args) > 1 {
		return "", ErrInvalid
	}

	key := ""

	if len(args) > 0 {
		key = args[0]
	}

	val, _ := self.buffMdl.Get(self.usr, key)
	return val, nil
}

func (self *Cmd) Set(args []string) error {
	if !self.logged || len(args) < 1 {
		return ErrInvalid
	}

	key := ""
	index := 0

	if len(args) > 1 {
		key = args[0]
		index = 1
	}

	return self.buffMdl.Set(self.usr, key, strings.Join(args[index:], " "))
}
