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
	// init config
	config := &configs.ServiceConfig{}

	if err := configs.Load([]configs.Config{config}); err != nil {
		panic(err)
	}

	// init caches
	baseCache := cache.BaseCache{"default"}
	bufferCache := &cache.MemoryCache{BaseCache: baseCache, Key: "buff"}
	saltCache := &cache.MemoryCache{BaseCache: baseCache, Key: "auth"}
	userCache := &cache.MemoryCache{BaseCache: baseCache, Key: "user"}

	if err := cache.Init([]cache.CacheInterface{
		bufferCache,
		saltCache,
		userCache,
	}); err != nil {
		panic(err)
	}

	port := ":" + strconv.Itoa(config.Port)
	l, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on " + port + "\n")

	defer l.Close()

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
