package store

import (
	"database/sql"
)

type Store interface {
	Init() error
	Insert(fieldMap map[string]interface{}) (sql.Result, error)
	Update(fieldMap map[string]interface{}, whereMap map[string]interface{}) (sql.Result, error)
	Query(selectFields []string, whereMap map[string]interface{}) (*sql.Rows, error)
	QueryRow(selectFields []string, whereMap map[string]interface{}) *sql.Row
}

func Init(stores []Store) error {
	for _, store := range stores {
		if err := store.Init(); err != nil {
			return err
		}
	}

	return nil
}
