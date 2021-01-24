package store

import (
	"database/sql"
	"fmt"
	"strings"
)

type BaseStore struct {
	DB       *sql.DB
	Table    string
	FieldMap map[string]string
}

func spread(length int, sep string) string {
	qmarks := make([]string, length)
	for i := range qmarks {
		qmarks[i] = "?"
	}
	return strings.Join(qmarks, sep)
}

func (store *BaseStore) Init() error {
	fieldList := []string{}
	for field, fieldType := range store.FieldMap {
		fieldList = append(fieldList, fmt.Sprintf("%s %s", field, fieldType))
	}

	fields := strings.Join(fieldList, ", ")
	cmd := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", store.Table, fields)

	if _, err := store.DB.Exec(cmd); err != nil {
		return err
	}

	return nil
}

func (store *BaseStore) Query(selectFields []string, whereMap map[string]interface{}) (*sql.Rows, error) {
	whereList := []string{}
	values := []interface{}{}

	for field, value := range whereMap {
		whereList = append(whereList, fmt.Sprintf("%s = ?", field))
		values = append(values, value)
	}

	fields := strings.Join(selectFields, ",")
	where := strings.Join(whereList, " AND ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", fields, store.Table, where)
	return store.DB.Query(query, values...)
}

func (store *BaseStore) QueryRow(selectFields []string, whereMap map[string]interface{}) *sql.Row {
	whereList := []string{}
	values := []interface{}{}

	for field, value := range whereMap {
		whereList = append(whereList, fmt.Sprintf("%s = ?", field))
		values = append(values, value)
	}

	fields := strings.Join(selectFields, ",")
	where := strings.Join(whereList, " AND ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", fields, store.Table, where)
	return store.DB.QueryRow(query, values...)
}

func (store *BaseStore) Insert(fieldMap map[string]interface{}) (sql.Result, error) {
	fieldList := []string{}
	values := []interface{}{}

	for field, value := range fieldMap {
		fieldList = append(fieldList, field)
		values = append(values, value)
	}

	fields := strings.Join(fieldList, ",")
	qmarks := spread(len(fieldList), ",")
	cmd := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", store.Table, fields, qmarks)
	return store.DB.Exec(cmd, values...)
}

func (store *BaseStore) Update(fieldMap map[string]interface{}, whereMap map[string]interface{}) (sql.Result, error) {
	updateList := []string{}
	whereList := []string{}
	values := []interface{}{}

	for field, value := range fieldMap {
		updateList = append(updateList, fmt.Sprintf("%s = ?", field))
		values = append(values, value)
	}

	for field, value := range whereMap {
		whereList = append(whereList, fmt.Sprintf("%s = ?", field))
		values = append(values, value)
	}

	update := strings.Join(updateList, ",")
	where := strings.Join(whereList, " AND ")
	cmd := fmt.Sprintf("UPDATE %s SET %s WHERE %s", store.Table, update, where)
	return store.DB.Exec(cmd, values...)
}
