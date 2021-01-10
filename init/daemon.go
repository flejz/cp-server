package main

import (
	"bufio"
	"fmt"
	"github.com/flejz/cp-server/configs"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const MIN = 1
const MAX = 100

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		c.Write([]byte(temp))

		result := strconv.Itoa(random()) + "\n"
		c.Write([]byte(string(result)))
	}
	c.Close()
}

func main() {
	config := &configs.DaemonConfig{}
	err := config.Load()

	if err != nil {
		panic(err)
	}

	l, err := net.Listen("tcp", ":2000")
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

		go handleConnection(conn)
	}
}
