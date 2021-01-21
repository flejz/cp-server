package db

import (
	"database/sql"
	"github.com/flejz/cp-server/configs"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type SQLiteDB struct {
	Config *configs.ServerConfig
	db     *sql.DB
}

func (s *SQLiteDB) Connect() error {
	var err error
	if _, err = os.Stat(s.Config.SQLitePath); os.IsNotExist(err) {
		file, err := os.Create(s.Config.SQLitePath)
		if err != nil {
			return err
		}
		file.Close()
	}

	if s.db, err = sql.Open("sqlite3", s.Config.SQLitePath); err != nil {
		return err
	}
	return nil
}

func (s *SQLiteDB) Base() *sql.DB {
	return s.db
}
