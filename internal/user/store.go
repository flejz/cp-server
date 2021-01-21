package user

import (
	"database/sql"
	"github.com/flejz/cp-server/internal/store"
)

type UserStore struct {
	store.BaseStore
}

func NewUserStore(db *sql.DB) store.Store {
	return &UserStore{
		store.BaseStore{
			db,
			"user",
			map[string]string{
				"usr":  "TEXT",
				"pwd":  "TEXT",
				"salt": "TEXT",
			},
		},
	}
}
