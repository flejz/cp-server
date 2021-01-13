package tcp

import (
	"crypto/sha256"
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/cache"
	err "github.com/flejz/cp-server/internal/error"
	"net"
	"strings"
)

type CmdHandler struct {
	authCache     cache.Cache
	bufferCache   cache.Cache
	conn          net.Conn
	serviceConfig *configs.ServiceConfig
}

func (cmdHandler *CmdHandler) Write(msg string) {
	cmdHandler.conn.Write([]byte(msg))
}

func (cmdHandler *CmdHandler) handle(c net.Conn, cmd string) error {
	parts := strings.Split(cmd, " ")
	action := parts[0]

	if action == "HELP" {
		return cmdHandler.help()
	} else if action == "LOGIN" {
		return cmdHandler.login(parts[1:])
	} else if action == "REG" {
		return cmdHandler.register(parts[1:])
	}

	return nil
}

func (cmdHandler *CmdHandler) help() error {
	msg := `LOGIN          logs you in
GET key        gets a value from a key
REG usr pwd    register a new user
SET key value  sets a new value to a key
`
	cmdHandler.Write(msg)
	return nil
}

func hash(value string, salt string) string {
	bytes := sha256.Sum256([]byte(fmt.Sprintf("%s.%s", value, salt)))
	return string(bytes[:])
}

func (cmdHandler *CmdHandler) register(creds []string) error {
	if len(creds) != 2 {
		return &err.InvalidCredentialsError{}
	}

	usr := creds[0]
	pwd := creds[1]

	pwdCheck, _ := cmdHandler.authCache.Get(usr)

	if pwdCheck != "" {
		return &err.UserExistsError{}
	}

	cmdHandler.authCache.Set(usr, hash(pwd, cmdHandler.serviceConfig.Salt))

	cmdHandler.Write("Registered\n")

	return nil
}

func (cmdHandler *CmdHandler) login(creds []string) error {
	if len(creds) != 2 {
		return &err.InvalidCredentialsError{}
	}

	usr := creds[0]
	pwd := creds[1]

	pwdCheck, pwdCheckErr := cmdHandler.authCache.Get(usr)

	if pwdCheckErr != nil || hash(pwd, cmdHandler.serviceConfig.Salt) != pwdCheck {
		return &err.InvalidCredentialsError{}
	}

	cmdHandler.Write("Logged\n")

	return nil
}
