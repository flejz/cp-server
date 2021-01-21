package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/proullon/ramsql/driver"
	"os"
)

func Connect() (*sql.DB, error) {
	config, err := Load()
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
