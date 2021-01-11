package tcp

import (
	"bufio"
	"fmt"
	"github.com/flejz/cp-server/internal/cache"
	"net"
	"strings"
)

type ConnHandler struct {
	cache cache.Cache
}

func (connHandler *ConnHandler) Handle(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	cmdHandler := &CmdHandler{cache: connHandler.cache}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		cmd := strings.TrimSpace(string(netData))
		if cmd == "STOP" || cmd == "EXIT" {
			break
		}

		cmdHandler.handle(c, cmd)
	}
	c.Close()
}
