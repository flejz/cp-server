package user

import (
	"database/sql"

	"github.com/flejz/cp-server/internal/repository"
)

type UserRepository struct {
	repository.BaseRepository
}

func NewUserRepository(db *sql.DB) repository.Repository {
	return &UserRepository{
		repository.BaseRepository{
			DB:    db,
			Table: "user",
			FieldMap: map[string]string{
				"usr":  "TEXT",
				"pwd":  "TEXT",
				"salt": "TEXT",
			},
		},
	}
}
