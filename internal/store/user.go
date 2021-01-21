package store

import (
	"database/sql"
)

type UserStore struct {
	BaseStore
}

func NewUserStore(db *sql.DB) StoreInterface {
	return &UserStore{
		BaseStore{
			db,
			"user",
			map[string]string{
				"usr": "TEXT NOT NULL PRIMARY KEY",
				"pwd": "TEXT",
			},
		},
	}
}

/*
func (store *UserStore) Init() error {
	createSQL := "CREATE TABLE IF NOT EXISTS " + store.table + " (usr TEXT NOT NULL PRIMARY KEY, pwd TEXT, salt TEXT)"
	if _, err := store.DB.Exec(createSQL); err != nil {
		return err
	}
	return nil
}
*/
