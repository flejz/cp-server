package store

import (
	"database/sql"
)

type BufferStore struct {
	BaseStore
}

func NewBufferStore(db *sql.DB) StoreInterface {
	return &BufferStore{
		BaseStore{
			db,
			"buffer",
			map[string]string{
				"usr":   "TEXT NOT NULL PRIMARY KEY",
				"key":   "TEXT",
				"value": "TEXT",
			},
		},
	}
}
