package model

import (
	"github.com/flejz/cp-server/configs"
	"github.com/flejz/cp-server/internal"
	"github.com/flejz/cp-server/internal/cache"
	err "github.com/flejz/cp-server/internal/error"
	"strings"
)

type Session struct {
	AuthCache     cache.Cache
	SaltCache     cache.Cache
	BufferCache   cache.Cache
	ServiceConfig *configs.ServiceConfig
	logged        bool
	usr           string
}

func (sess *Session) Register(usr, pwd string) error {
	if pass, _ := sess.AuthCache.Get(usr); pass != "" {
		return &err.InvalidError{}
	}

	salt := util.Salt()

	if err := sess.SaltCache.Set(usr, salt); err != nil {
		return err
	}

	if err := sess.AuthCache.Set(usr, util.Hash(pwd, salt)); err != nil {
		return err
	}

	return nil
}

func (sess *Session) Login(usr, pwd string) error {
	salt, saltErr := sess.SaltCache.Get(usr)
	pass, passErr := sess.AuthCache.Get(usr)

	if saltErr != nil ||
		passErr != nil ||
		util.Hash(pwd, salt) != pass {
		return &err.InvalidError{}
	}

	sess.logged = true
	sess.usr = usr

	return nil
}

func (sess *Session) Logout() error {
	sess.logged = false
	sess.usr = ""

	return nil
}

func (sess *Session) Get() (string, error) {
	if !sess.logged {
		return "", &err.InvalidError{}
	}

	val, _ := sess.BufferCache.Get(sess.usr)
	return val, nil
}

func (sess *Session) Set(vals []string) error {
	if !sess.logged || len(vals) < 1 {
		return &err.InvalidError{}
	}

	return sess.BufferCache.Set(sess.usr, strings.Join(vals, " "))
}
