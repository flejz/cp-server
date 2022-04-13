package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/user"
)

type ConnHandler struct {
	buffMdl buffer.BufferModel
	usrMdl  user.UserModel
}

func (c *ConnHandler) Handle(conn net.Conn) {
	fmt.Printf("serving %s\n", conn.RemoteAddr().String())
	cmd := &buffer.Cmd{
		BuffMdl: c.buffMdl,
		UsrMdl:  c.usrMdl,
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
