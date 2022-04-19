package tcp

import (
	"net"
	"strings"

	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/user"
)

type CPController struct {
	conn          net.Conn
	session       CPSession
	bufferService *buffer.BufferService
	usrService    *user.UserService
}

func (c *CPController) Write(msg string) {
	c.conn.Write([]byte(msg))
}

func (c *CPController) help() error {
	c.Write("LOGIN  usr pwd\n" +
		"LOGOUT\n" +
		"REG    usr pwd\n" +
		"GET    [key]\n" +
		"SET    [key] val\n" +
		"EXIT\n")
	return nil
}

func (c *CPController) login(args []string) error {
	if len(args) != 2 {
		return ErrInvalid
	}
	if err := c.session.Login(args[0], args[1]); err != nil {
		return err
	}

	c.Write("logged\n")
	return nil
}

func (c *CPController) logout() error {
	if err := c.session.Logout(); err != nil {
		return err
	}
	c.Write("logged out\n")
	return nil
}

func (c *CPController) register(args []string) error {
	if len(args) != 2 {
		return ErrInvalid
	}
	if err := c.session.Register(args[0], args[1]); err != nil {
		return err
	}

	c.Write("registered\n")
	return nil
}

func (c *CPController) get(args []string) error {
	if !c.session.logged || len(args) > 1 {
		return ErrInvalid
	}

	key := ""

	if len(args) > 0 {
		key = args[0]
	}

	val, err := c.bufferService.Get(c.session.usr, key)

	if err != nil {
		return err
	}

	if val != "" {
		c.Write(val + "\n")
	}

	return nil
}

func (c *CPController) set(args []string) error {
	if !c.session.logged || len(args) < 1 {
		return ErrInvalid
	}

	key := ""
	index := 0

	if len(args) > 1 {
		key = args[0]
		index = 1
	}

	return c.bufferService.Set(c.session.usr, key, strings.Join(args[index:], " "))
}
