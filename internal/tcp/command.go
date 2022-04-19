package tcp

import (
	"bufio"
	"net"
	"strings"
)

type CPCommandHandler struct {
	conn       net.Conn
	controller CPController
}

func (c *CPCommandHandler) Parse() error {
	data, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		return err
	}

	raw := strings.TrimSpace(string(data))

	cmdArgs := strings.Split(raw, " ")
	action := strings.ToUpper(cmdArgs[0])
	args := cmdArgs[1:]

	switch action {
	case "HELP":
		return c.controller.help()
	case "LOGIN", "LOG":
		return c.controller.login(args)
	case "LOGOUT":
		return c.controller.logout()
	case "REG":
		return c.controller.register(args)
	case "SET":
		return c.controller.set(args)
	case "GET":
		return c.controller.get(args)
	case "EXIT", "^C":
		return ErrInterrupted
	default:
		return ErrInvalid
	}

}
