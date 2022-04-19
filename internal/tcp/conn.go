package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/user"
)

type TCPConnHandler struct {
	bufferService buffer.BufferService
	usrService    user.UserService
}

func (c *TCPConnHandler) Handle(conn net.Conn) {
	fmt.Printf("serving %s\n", conn.RemoteAddr().String())
	cmd := &buffer.Cmd{
		BufferService: c.bufferService,
		UsrService:    c.usrService,
	}

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}

		raw := strings.TrimSpace(string(data))
		args := strings.Split(raw, " ")
		action := strings.ToUpper(args[0])

		if err := Parse(conn, cmd, action, args[1:]); err != nil {
			switch err {
			case ErrInterrupted:
				conn.Close()
				fmt.Printf("closing %s\n", conn.RemoteAddr().String())
				return
			default:
				Write(conn, fmt.Errorf("%v", err).Error())
			}
		}
	}
}
