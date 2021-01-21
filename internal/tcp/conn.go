package tcp

import (
	"bufio"
	"fmt"
	"github.com/flejz/cp-server/internal/errors"
	"github.com/flejz/cp-server/internal/model"
	"net"
	"strings"
)

type ConnHandler struct {
	UserModel   model.User
	BufferModel model.Buffer
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
			fmt.Println(err)
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
				Write(conn, fmt.Sprintf("%s\n", err.Error()))
				Write(conn, err.(*errors.Error).ErrorStack())
			}
		}
	}
}
