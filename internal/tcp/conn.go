package tcp

import (
	"bufio"
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/cache"
	"net"
	"strings"
)

type ConnHandler struct {
	AuthCache     cache.Cache
	BufferCache   cache.Cache
	ServiceConfig *configs.ServiceConfig
}

func (connHandler *ConnHandler) Handle(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	cmdHandler := &CmdHandler{
		authCache:     connHandler.AuthCache,
		bufferCache:   connHandler.BufferCache,
		conn:          conn,
		serviceConfig: connHandler.ServiceConfig,
	}

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		cmd := strings.TrimSpace(string(data))
		if cmd == "STOP" || cmd == "EXIT" {
			break
		}

		cmdErr := cmdHandler.handle(conn, cmd)
		if cmdErr != nil {
			cmdHandler.Write(fmt.Sprintf("%s\n", cmdErr.Error()))
		}
	}
	conn.Close()
}
