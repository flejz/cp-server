package cache

import (
	"github.com/flejz/cp-server/internal/errors"
	"github.com/flejz/cp-server/internal/store"
)

type SQLiteCache struct {
	Store      store.StoreInterface
	DefaultKey string
}

func (sqlite *SQLiteCache) Init() error {
	return nil
}

func (sqlite *SQLiteCache) key(key string) string {
	if key == "" {
		return sqlite.DefaultKey
	} else {
		return key
	}
}

func (sqlite *SQLiteCache) Get(usr, key string) (string, error) {
	if usr == "" {
		return "", &errors.KeyNotSetError{}
	}

	selectFields := []string{"value"}
	whereMap := map[string]interface{}{
		"usr": usr,
		"key": sqlite.key(key),
	}

	row, err := sqlite.Store.Query(selectFields, whereMap)
	if err != nil {
		return "", nil
	}

	for row.Next() {
		var value string
		row.Scan(&value)
		return value, nil
	}

	return "", &errors.NotFoundStoreError{}
}

func (sqlite *SQLiteCache) Set(usr, key, value string) error {
	if usr == "" {
		return &errors.KeyNotSetError{}
	}

	if value == "" {
		return &errors.ValueNotSetError{}
	}

	fieldMap := map[string]interface{}{
		"usr":   usr,
		"key":   sqlite.key(key),
		"value": value,
	}

	value, err := sqlite.Get(usr, key)
	switch err.(type) {
	case *errors.NotFoundStoreError:
		_, err := sqlite.Store.Insert(fieldMap)
		return err
	default:
		if err != nil {
			return err
		}
	}

	fieldMap = map[string]interface{}{
		"value": value,
	}
	whereMap := map[string]interface{}{
		"usr": usr,
		"key": sqlite.key(key),
	}

	_, err = sqlite.Store.Update(fieldMap, whereMap)
	return err
}
