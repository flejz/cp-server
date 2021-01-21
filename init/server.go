package main

import (
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/cache"
	"github.com/flejz/cp-server/internal/db"
	"github.com/flejz/cp-server/internal/model"
	"github.com/flejz/cp-server/internal/store"
	"github.com/flejz/cp-server/internal/tcp"
	"log"
	"net"
	"strconv"
)

func main() {
	// init config
	config := &configs.ServerConfig{}

	if err := configs.Load([]configs.Config{config}); err != nil {
		panic(err)
	}

	// init database
	sqliteDB := &db.SQLiteDB{Config: config}
	defer sqliteDB.Base().Close()

	if err := db.Connect([]db.DB{sqliteDB}); err != nil {
		panic(err)
	}

	// init stores
	bufferStore := store.NewBufferStore(sqliteDB.Base())
	saltStore := store.NewSaltStore(sqliteDB.Base())
	userStore := store.NewUserStore(sqliteDB.Base())

	if err := store.Init([]store.StoreInterface{
		bufferStore,
		saltStore,
		userStore,
	}); err != nil {
		panic(err)
	}

	// init caches
	bufferCache := &cache.SQLiteCache{bufferStore, "default"}
	saltCache := &cache.SQLiteCache{saltStore, "default"}
	userCache := &cache.SQLiteCache{userStore, "default"}

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
