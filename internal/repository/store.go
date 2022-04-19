package repository

import (
	"database/sql"
)

type Repository interface {
	Init() error
	Insert(fieldMap map[string]interface{}) (sql.Result, error)
	Update(fieldMap map[string]interface{}, whereMap map[string]interface{}) (sql.Result, error)
	Query(selectFields []string, whereMap map[string]interface{}) (*sql.Rows, error)
	QueryRow(selectFields []string, whereMap map[string]interface{}) *sql.Row
}

func Init(repositorys []Repository) error {
	for _, repository := range repositorys {
		if err := repository.Init(); err != nil {
			return err
		}
	}

	return nil
}
