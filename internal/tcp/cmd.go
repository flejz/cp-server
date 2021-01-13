package tcp

import (
	"crypto/sha256"
	"fmt"
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal/cache"
	err "github.com/flejz/cp-server/internal/error"
	"math/rand"
	"net"
	"strings"
	"time"
)

func hash(value string, salt string) string {
	bytes := sha256.Sum256([]byte(fmt.Sprintf("%s.%s", value, salt)))
	return string(bytes[:])
}

func salt() string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder

	charSet := []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP0123456789")
	length := 20
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}
	return output.String()
}

type Mach struct {
	authCache     cache.Cache
	saltCache     cache.Cache
	bufferCache   cache.Cache
	conn          net.Conn
	serviceConfig *configs.ServiceConfig
	logged        bool
	usr           string
}

func (mach *Mach) Write(msg string) {
	mach.conn.Write([]byte(msg))
}

func (mach *Mach) handle(c net.Conn, cmd string) error {
	parts := strings.Split(cmd, " ")
	action := parts[0]

	var err error

	if action == "HELP" {
		err = mach.help()
	} else if action == "LOGIN" {
		err = mach.login(parts[1:])
	} else if action == "LOGOUT" {
		err = mach.logout()
	} else if action == "REG" {
		err = mach.register(parts[1:])
	} else if action == "SET" {
		err = mach.set(parts[1:])
	} else if action == "GET" {
		err = mach.get()
	}

	return err
}

func (mach *Mach) help() error {
	msg := `LOGIN  usr pwd    logs you in
LOGOUT            gets a value from a key
GET               gets a value from a key
REG    usr pwd    register a new user
SET    value      sets a new value to a key
`
	mach.Write(msg)
	return nil
}

func (mach *Mach) register(creds []string) error {
	if len(creds) != 2 {
		return &err.InvalidError{}
	}

	usr := creds[0]
	pwd := creds[1]

	pwdCheck, _ := mach.authCache.Get(usr)

	if pwdCheck != "" {
		return &err.ExistsError{}
	}

	salt := salt()
	mach.saltCache.Set(usr, salt)
	mach.authCache.Set(usr, hash(pwd, salt))

	mach.Write("Registered\n")

	return nil
}

func (mach *Mach) login(creds []string) error {
	if len(creds) != 2 {
		return &err.InvalidError{}
	}

	usr := creds[0]
	pwd := creds[1]

	salt, saltErr := mach.saltCache.Get(usr)
	pass, passErr := mach.authCache.Get(usr)

	if saltErr != nil ||
		passErr != nil ||
		hash(pwd, salt) != pass {
		return &err.InvalidError{}
	}

	mach.logged = true
	mach.usr = usr
	mach.Write("Logged\n")

	return nil
}

func (mach *Mach) logout() error {
	mach.logged = false
	mach.usr = ""

	return nil
}

func (mach *Mach) get() error {
	if !mach.logged {
		return &err.InvalidError{}
	}

	value, err := mach.bufferCache.Get(mach.usr)
	if value != "" {
		mach.Write(value + "\n")
	}
	return err
}

func (mach *Mach) set(vals []string) error {
	if !mach.logged || len(vals) < 1 {
		return &err.InvalidError{}
	}

	return mach.bufferCache.Set(mach.usr, strings.Join(vals, " "))
}
