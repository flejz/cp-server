package store

import (
	"database/sql"
)

type SaltStore struct {
	BaseStore
}

func NewSaltStore(db *sql.DB) StoreInterface {
	return &SaltStore{
		BaseStore{
			db,
			"salt",
			map[string]string{
				"usr":  "TEXT NOT NULL PRIMARY KEY",
				"salt": "TEXT",
			},
		},
	}
}

/*
func (store *SaltStore) Init() error {
	createSQL := "CREATE TABLE IF NOT EXISTS " + store.table + " (usr TEXT NOT NULL PRIMARY KEY, pwd TEXT, salt TEXT)"
	if _, err := store.DB.Exec(createSQL); err != nil {
		return err
	}
	return nil
}
*/
