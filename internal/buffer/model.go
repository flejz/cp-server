package buffer

import (
	"github.com/flejz/cp-server/internal/store"
)

type BufferModel struct {
	Store store.Store
}

func (self BufferModel) Get(usr, key string) (string, error) {
	fieldList := []string{"value"}
	whereMap := map[string]interface{}{
		"usr":   usr,
		"keyid": key,
	}

	row := self.Store.QueryRow(fieldList, whereMap)

	if err := row.Err(); err != nil {
		return "", err
	}

	var value string
	if err := row.Scan(&value); err != nil {
		return "", err
	}

	return value, nil
}

func (self BufferModel) Set(usr, key, value string) error {
	fieldMap := map[string]interface{}{
		"usr":   usr,
		"keyid": key,
		"value": value,
	}

	if _, err := self.Store.Insert(fieldMap); err != nil {
		return err
	}

	return nil
}
