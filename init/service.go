package main

import (
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/tcp"
	"log"
	"net"
	"strconv"
)

const MIN = 1
const MAX = 100

func main() {
	config, configErr := configs.LoadServiceConfig()

	if configErr != nil {
		panic(configErr)
	}

	port := ":" + strconv.Itoa(config.Port)
	l, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on " + port + "\n")

	defer l.Close()
	connHandler := &tcp.ConnHandler{}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go connHandler.Handle(conn)
	}
}
