package tcp

import (
	"github.com/flejz/cp-server/internal/cache"
	"net"
	"strings"
)

type CmdHandler struct {
	c     net.Conn
	cache cache.Cache
}

func (cmdHandler *CmdHandler) write(msg string) {
	cmdHandler.c.Write([]byte(msg))
}

func (cmdHandler *CmdHandler) handle(c net.Conn, cmd string) error {
	parts := strings.Split(cmd, " ")
	action := parts[0]

	if action == "HELP" {
		return cmdHandler.help()
	} else if action == "LOGIN" {
		return cmdHandler.login(parts[1], parts[2])
	}

	return nil
}

func (cmdHandler *CmdHandler) help() error {
	msg := `
LOGIN			logs you in
SET	key value	sets a new value to a key
GET	key		gets a value from a key
	`
	cmdHandler.write(msg)
	return nil
}

func (cmdHandler *CmdHandler) login(usr, pwd string) error {
	cmdHandler.write("Login successfull")
	return nil
}
