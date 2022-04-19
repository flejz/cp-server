package buffer

import (
	"database/sql"

	"github.com/flejz/cp-server/internal/repository"
)

type BufferRepository struct {
	repository.BaseRepository
}

func NewBufferRepository(db *sql.DB) repository.Repository {
	return &BufferRepository{
		repository.BaseRepository{
			DB:    db,
			Table: "buffer",
			FieldMap: map[string]string{
				"usr":   "TEXT",
				"keyid": "TEXT",
				"value": "TEXT",
			},
		},
	}
}
