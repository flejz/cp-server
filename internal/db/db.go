package db

import (
	"database/sql"
	"os"

	"github.com/flejz/cp-server/internal/config"
	_ "github.com/mattn/go-sqlite3" // just the import is needed here in order to make it work
	_ "github.com/proullon/ramsql/driver"
)

// Open the db connection
func Open(cfg *config.Config) (*sql.DB, error) {
	switch cfg.DB.Mode {
	case config.MemoryMode:
		return sql.Open("ramsql", cfg.DB.Name)
	case config.SQLiteMode:
		if _, err := os.Stat(cfg.DB.Path); os.IsNotExist(err) {
			file, err := os.Create(cfg.DB.Path)
			if err != nil {
				return nil, err
			}
			file.Close()
		}

		return sql.Open("sqlite3", cfg.DB.Path)
	}
	return nil, ErrInvalidDBMode
}
