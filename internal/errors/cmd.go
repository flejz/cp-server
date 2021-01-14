package errors

import (
	"fmt"
)

type ExistsError struct{}

func (e *ExistsError) Error() string {
	return fmt.Sprint("Already exists")
}

type InvalidError struct{}

func (e *InvalidError) Error() string {
	return fmt.Sprint("Invalid")
}
