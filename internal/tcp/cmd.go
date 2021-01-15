package tcp

import (
	"github.com/flejz/cp-server/internal/errors"
	"net"
)

func Write(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
}

func Parse(conn net.Conn, sess *Session, action string, args []string) error {
	switch action {
	case "EXIT":
		return &errors.InterruptionError{}
	case "HELP":
		Write(conn, "GET\n"+
			"EXIT\n"+
			"LOGIN  usr pwd\n"+
			"LOGOUT\n"+
			"REG    usr pwd\n"+
			"SET    val\n")
	case "LOGIN":
		if len(args) != 2 {
			return &errors.InvalidError{}
		}
		if err := sess.Login(args[0], args[1]); err != nil {
			return err
		}

		Write(conn, "Logged\n")
	case "LOGOUT":
		return sess.Logout()
	case "REG":
		if len(args) != 2 {
			return &errors.InvalidError{}
		}
		if err := sess.Register(args[0], args[1]); err != nil {
			return err
		}

		Write(conn, "Registered\n")
	case "SET":
		return sess.Set(args)
	case "GET":
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
