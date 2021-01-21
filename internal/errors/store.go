package errors

import (
	"fmt"
)

type NotFoundStoreError struct{}

func (e *NotFoundStoreError) Error() string {
	return fmt.Sprint("Any record found")
}
