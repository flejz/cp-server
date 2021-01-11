package main

import (
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/tcp"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

const MIN = 1
const MAX = 100

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func main() {
	config, configErr := configs.LoadServiceConfig()

	if configErr != nil {
		panic(configErr)
	}

	port := ":" + strconv.Itoa(config.Port)
	l, err := net.Listen("tcp", port)

	fmt.Printf("Listening on " + port + "\n")

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go tcp.HandleConnection(conn)
	}
}
