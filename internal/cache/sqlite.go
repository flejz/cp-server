package cache

import (
	"fmt"
	"github.com/flejz/cp-server/internal/errors"
	"github.com/flejz/cp-server/internal/store"
)

type SQLiteCache struct {
	BaseCache
	Store store.StoreInterface
}

func (sqlite *SQLiteCache) Init() error {
	return nil
}

func (sqlite *SQLiteCache) Get(usr, key string) (string, error) {
	fmt.Print(">>> GET\n")
	fmt.Printf(">>> %s\n", usr)
	fmt.Printf(">>> %s\n", sqlite.key(key))
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
	fmt.Print(">>> SET\n")
	fmt.Printf(">>> %s\n", usr)
	fmt.Printf(">>> %s\n", sqlite.key(key))
	fmt.Printf(">>> %s\n", value)
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
