package tcp

import (
	"net"
	"strings"
)

func HandleCommand(c net.Conn, command string) {
	cmd := strings.Trim(command, " ")

	if cmd == "HELP" {
		handleHelpCommand(c)
	}
}

func writeToConn(c net.Conn, msg string) {
	c.Write([]byte(msg))
}

func handleHelpCommand(c net.Conn) {
	msg := `
LOGIN              login you in
SET key value      sets a new value to a key
SET key            gets a value from a key
	`

	writeToConn(c, msg)
}
