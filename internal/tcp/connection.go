package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		command := strings.TrimSpace(string(netData))
		if command == "STOP" || command == "EXIT" {
			break
		}

		HandleCommand(c, command)
	}
	c.Close()
}
