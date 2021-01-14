package main

import (
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/cache"
	"github.com/flejz/cp-server/internal/model"
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

	userCache := &cache.Memory{Key: "auth"}
	userCache.Init()

	saltCache := &cache.Memory{Key: "auth"}
	saltCache.Init()

	bufferCache := &cache.Memory{Key: "buff"}
	bufferCache.Init()

	saltModel := model.Salt{Cache: saltCache}
	userModel := model.User{Cache: userCache, SaltModel: saltModel}
	bufferModel := model.Buffer{Cache: bufferCache}

	connHandler := tcp.ConnHandler{
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
