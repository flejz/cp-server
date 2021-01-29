package tcp

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/config"
	"github.com/flejz/cp-server/internal/db"
	"github.com/flejz/cp-server/internal/store"
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

	// init stores
	buffStore := buffer.NewBufferStore(db)
	usrStore := user.NewUserStore(db)

	if err := store.Init([]store.Store{buffStore, usrStore}); err != nil {
		panic(err)
	}

	// init models
	buffMdl := buffer.BufferModel{Store: buffStore}
	usrMdl := user.UserModel{Store: usrStore}

	// getting proper ports
	port := ":" + strconv.Itoa(cfg.TCP.Port)
	l, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	fmt.Printf("listening on " + port + "\n")

	connHandler := ConnHandler{
		buffMdl,
		usrMdl,
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go connHandler.Handle(conn)
	}
}
