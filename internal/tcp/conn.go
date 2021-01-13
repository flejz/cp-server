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
	SaltCache     cache.Cache
	BufferCache   cache.Cache
	ServiceConfig *configs.ServiceConfig
}

func (connHandler *ConnHandler) Handle(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	mach := &Mach{
		authCache:     connHandler.AuthCache,
		saltCache:     connHandler.SaltCache,
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

		cmdErr := mach.handle(conn, cmd)
		if cmdErr != nil {
			mach.Write(fmt.Sprintf("%s\n", cmdErr.Error()))
		}
	}
	fmt.Printf("Closing %s\n", conn.RemoteAddr().String())
	conn.Close()
}
