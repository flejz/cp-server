package tcp

import (
	"github.com/flejz/cp-server/internal/errors"
	"github.com/flejz/cp-server/internal/model"
	"net"
	"strings"
)

func Write(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
}

func ParseCmd(conn net.Conn, sess *model.Session, cmd string) error {
	parts := strings.Split(cmd, " ")
	action := strings.ToUpper(parts[0])

	if action == "HELP" {
		Write(conn, `GET
LOGIN  usr pwd
LOGOUT
REG    usr pwd
SET    val
`)
	} else if action == "LOGIN" {
		if len(parts) != 3 {
			return &errors.InvalidError{}
		}
		if err := sess.Login(parts[1], parts[2]); err != nil {
			return err
		}

		Write(conn, "Logged\n")
	} else if action == "LOGOUT" {
		return sess.Logout()
	} else if action == "REG" {
		if len(parts) != 3 {
			return &errors.InvalidError{}
		}
		if err := sess.Register(parts[1], parts[2]); err != nil {
			return err
		}

		Write(conn, "Registered\n")
	} else if action == "SET" {
		return sess.Set(parts[1:])
	} else if action == "GET" {
		val, err := sess.Get()

		if err != nil {
			return err
		}

		if val != "" {
			Write(conn, val+"\n")
		}
	}

	return nil
}
