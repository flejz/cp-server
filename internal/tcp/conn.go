package tcp

import (
	"bufio"
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/cache"
	"github.com/flejz/cp-server/internal/model"
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
	sess := &model.Session{
		AuthCache:     connHandler.AuthCache,
		SaltCache:     connHandler.SaltCache,
		BufferCache:   connHandler.BufferCache,
		ServiceConfig: connHandler.ServiceConfig,
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

		if err := ParseCmd(conn, sess, cmd); err != nil {
			Write(conn, fmt.Sprintf("%s\n", err.Error()))
		}
	}
	fmt.Printf("Closing %s\n", conn.RemoteAddr().String())
	conn.Close()
}
