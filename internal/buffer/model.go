package buffer

import (
	"database/sql"
	"github.com/flejz/cp-server/internal/store"
	"strings"
)

type BufferModel struct {
	Store store.Store
}

func keyid(key string) string {
	if strings.Trim(key, " ") == "" {
		return "_default"
	} else {
		return key
	}
}

func (self BufferModel) Get(usr, key string) (string, error) {
	fieldList := []string{"value"}
	whereMap := map[string]interface{}{
		"usr":   usr,
		"keyid": keyid(key),
	}

	row := self.Store.QueryRow(fieldList, whereMap)

	if err := row.Err(); err != nil {
		return "", err
	}

	var value string
	if err := row.Scan(&value); err != nil {
		switch err {
		case sql.ErrNoRows:
			return "", nil
		default:
			return "", err
		}
	}

	return value, nil
}

func (self BufferModel) Set(usr, key, value string) error {
	cachedValue, err := self.Get(usr, key)
	if err != nil {
		return err
	}
	if cachedValue != "" {
		fieldMap := map[string]interface{}{
			"value": value,
		}
		whereMap := map[string]interface{}{
			"usr":   usr,
			"keyid": keyid(key),
		}

		if _, err := self.Store.Update(fieldMap, whereMap); err != nil {
			return err
		}

		return nil
	}
	fieldMap := map[string]interface{}{
		"usr":   usr,
		"keyid": keyid(key),
		"value": value,
	}

	if _, err := self.Store.Insert(fieldMap); err != nil {
		return err
	}

	return nil
}
