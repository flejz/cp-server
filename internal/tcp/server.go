package tcp

import (
	"fmt"
	"github.com/flejz/cp-server/internal/buffer"
	"github.com/flejz/cp-server/internal/db"
	"github.com/flejz/cp-server/internal/store"
	"github.com/flejz/cp-server/internal/user"
	"log"
	"net"
	"strconv"
)

func Listen() {
	config, err := Load()
	if err != nil {
		panic(err)
	}

	// init db
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// init stores
	bufferStore := buffer.NewBufferStore(db)
	userStore := user.NewUserStore(db)

	if err := store.Init([]store.Store{bufferStore, userStore}); err != nil {
		panic(err)
	}

	// init models
	bufferModel := buffer.BufferModel{bufferStore}
	userModel := user.UserModel{userStore}

	// getting proper ports
	port := ":" + strconv.Itoa(config.Port)
	l, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	fmt.Printf("Listening on " + port + "\n")

	connHandler := ConnHandler{
		UserModel:   userModel,
		BufferModel: bufferModel,
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go connHandler.Handle(conn)
	}
}
