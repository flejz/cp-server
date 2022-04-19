package buffer

import (
	"database/sql"
	"strings"

	"github.com/flejz/cp-server/internal/repository"
)

type BufferService struct {
	Repository repository.Repository
}

func keyid(key string) string {
	if strings.Trim(key, " ") == "" {
		return "_default"
	} else {
		return key
	}
}

func (self BufferService) Get(usr, key string) (string, error) {
	fieldList := []string{"value"}
	whereMap := map[string]interface{}{
		"usr":   usr,
		"keyid": keyid(key),
	}

	row := self.Repository.QueryRow(fieldList, whereMap)

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

func (self BufferService) Set(usr, key, value string) error {
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

		if _, err := self.Repository.Update(fieldMap, whereMap); err != nil {
			return err
		}

		return nil
	}
	fieldMap := map[string]interface{}{
		"usr":   usr,
		"keyid": keyid(key),
		"value": value,
	}

	if _, err := self.Repository.Insert(fieldMap); err != nil {
		return err
	}

	return nil
}
