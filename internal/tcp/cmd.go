package tcp

import (
	"net"

	"github.com/flejz/cp-server/internal/buffer"
)

func Write(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
}

func Parse(conn net.Conn, cmd *buffer.Cmd, action string, args []string) error {
	switch action {
	case "EXIT":
		return ErrInterrupted
	case "HELP":
		Write(conn, "GET\n"+
			"EXIT\n"+
			"LOGIN  usr pwd\n"+
			"LOGOUT\n"+
			"REG    usr pwd\n"+
			"SET    val\n")
	case "LOGIN", "LOG":
		if len(args) != 2 {
			return ErrInvalid
		}
		if err := cmd.Login(args[0], args[1]); err != nil {
			return err
		}

		Write(conn, "logged\n")
	case "LOGOUT":
		if err := cmd.Logout(); err != nil {
			return err
		}
		Write(conn, "logged out\n")
	case "REG":
		if len(args) != 2 {
			return ErrInvalid
		}
		if err := cmd.Register(args[0], args[1]); err != nil {
			return err
		}

		Write(conn, "registered\n")
	case "SET":
		return cmd.Set(args)
	case "GET":
		val, err := cmd.Get(args)

		if err != nil {
			return err
		}

		if val != "" {
			Write(conn, val+"\n")
		}
	}

	return nil
}
