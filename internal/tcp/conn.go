package tcp

import (
	"bufio"
	"fmt"
	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/errors"
	"github.com/flejz/cp-server/internal/user"
	"net"
	"strings"
)

type ConnHandler struct {
	UserModel   user.UserModel
	BufferModel buffer.BufferModel
}

func (connHandler *ConnHandler) Handle(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	sess := &Session{
		UserModel:   connHandler.UserModel,
		BufferModel: connHandler.BufferModel,
	}

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}

		cmd := strings.TrimSpace(string(data))
		args := strings.Split(cmd, " ")
		action := strings.ToUpper(args[0])

		if err := Parse(conn, sess, action, args[1:]); err != nil {
			switch err.(type) {
			case *errors.InterruptionError:
				conn.Close()
				fmt.Printf("Closing %s\n", conn.RemoteAddr().String())
				return
			default:
				fmt.Printf(">>> dang %+v %t\n", err, err)
				Write(conn, fmt.Errorf("%v\n", err).Error())
			}
		}
	}
}
