package tcp

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/config"
	"github.com/flejz/cp-server/internal/db"
	"github.com/flejz/cp-server/internal/repository"
	"github.com/flejz/cp-server/internal/user"
)

// Listen starts the tcp server
func Listen() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	// init db
	db, err := db.Open(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// init repositorys
	bufferRepository := buffer.NewBufferRepository(db)
	usrRepository := user.NewUserRepository(db)

	if err := repository.Init([]repository.Repository{bufferRepository, usrRepository}); err != nil {
		panic(err)
	}

	// init models
	bufferService := &buffer.BufferService{Repository: &bufferRepository}
	usrService := &user.UserService{Repository: &usrRepository}

	// getting proper ports
	port := ":" + strconv.Itoa(cfg.TCP.Port)
	socket, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	defer socket.Close()

	fmt.Printf("listening on " + port + "\n")

	connHandler := TCPConnHandler{
		bufferService,
		usrService,
	}

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go connHandler.Handle(conn)
	}
}
