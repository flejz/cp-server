package errors

import (
	"fmt"
)

type KeyNotFoundError struct {
	Key string
}

func (k *KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key was not found: %s", k.Key)
}

type KeyNotSetError struct{}

func (k *KeyNotSetError) Error() string {
	return fmt.Sprint("Key not set")
}

type ValueNotSetError struct{}

func (k *ValueNotSetError) Error() string {
	return fmt.Sprintf("Value not set")
}
