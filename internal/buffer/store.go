package buffer

import (
	"database/sql"
	"github.com/flejz/cp-server/internal/store"
)

type BufferStore struct {
	store.BaseStore
}

func NewBufferStore(db *sql.DB) store.Store {
	return &BufferStore{
		store.BaseStore{
			db,
			"buffer",
			map[string]string{
				"usr":   "TEXT",
				"keyid": "TEXT",
				"value": "TEXT",
			},
		},
	}
}
