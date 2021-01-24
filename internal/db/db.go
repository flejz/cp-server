package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/proullon/ramsql/driver"
)

func Connect(defaults bool) (*sql.DB, error) {
	config, err := Load(defaults)
	if err != nil {
		return nil, err
	}

	switch config.Type {
	case Memory:
		return sql.Open("ramsql", config.MemoryName)
	case SQLite:
		if _, err := os.Stat(config.SQLitePath); os.IsNotExist(err) {
			file, err := os.Create(config.SQLitePath)
			if err != nil {
				return nil, err
			}
			file.Close()
		}

		return sql.Open("sqlite3", config.SQLitePath)
	}
	return nil, nil
}
