package main

import (
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/cache"
	"github.com/flejz/cp-server/internal/tcp"
	"log"
	"net"
	"strconv"
)

func main() {
	serviceConfig, serviceConfigErr := configs.Load()

	if serviceConfigErr != nil {
		panic(serviceConfigErr)
	}

	port := ":" + strconv.Itoa(serviceConfig.Port)
	l, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on " + port + "\n")

	defer l.Close()

	authCache := &cache.Memory{Key: "auth"}
	authCache.Init()

	saltCache := &cache.Memory{Key: "auth"}
	saltCache.Init()

	bufferCache := &cache.Memory{Key: "buff"}
	bufferCache.Init()

	connHandler := &tcp.ConnHandler{
		AuthCache:     authCache,
		SaltCache:     saltCache,
		BufferCache:   bufferCache,
		ServiceConfig: serviceConfig,
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go connHandler.Handle(conn)
	}
}
