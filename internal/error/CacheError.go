package error

import (
	"fmt"
)

type KeyNotFoundError struct {
	Key string
}

func (k *KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key was not found: %s", k.Key)
}
