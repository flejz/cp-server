package tcp

import (
	"bufio"
	"fmt"
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
	sess := &model.Session{
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
		if strings.ToUpper(cmd) == "STOP" || strings.ToUpper(cmd) == "EXIT" {
			break
		}

		if err := ParseCmd(conn, sess, cmd); err != nil {
			Write(conn, fmt.Sprintf("%s\n", err.Error()))
		}
	}
	fmt.Printf("Closing %s\n", conn.RemoteAddr().String())
	conn.Close()
}
