package tcp

import (
	"fmt"
	"net"

	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/user"
)

type TCPConnHandler struct {
	bufferService *buffer.BufferService
	usrService    *user.UserService
}

func (c *TCPConnHandler) Handle(conn net.Conn) {
	fmt.Printf("serving %s\n", conn.RemoteAddr().String())

	session := CPSession{
		usrService: c.usrService,
	}

	controller := CPController{
		conn,
		session,
		c.bufferService,
		c.usrService,
	}

	handler := CPCommandHandler{
		conn,
		controller,
	}

	for {

		if err := handler.Parse(); err != nil {
			switch err {
			case ErrInterrupted:
				conn.Close()
				fmt.Printf("closing %s\n", conn.RemoteAddr().String())
				return
			default:
				controller.Write(fmt.Errorf("%v\n", err).Error())
			}
		}
	}
}
