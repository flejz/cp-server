package tcp

import (
	"github.com/flejz/cp-server/internal/user"
)

type CPSession struct {
	usrService *user.UserService
	logged     bool
	usr        string
}

func (s *CPSession) Register(usr, pwd string) error {
	return s.usrService.Register(usr, pwd)
}

func (s *CPSession) Login(usr, pwd string) error {
	if err := s.usrService.Validate(usr, pwd); err != nil {
		return err
	}

	s.logged = true
	s.usr = usr

	return nil
}

func (s *CPSession) Logout() error {
	s.logged = false
	s.usr = ""

	return nil
}
