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

type InterruptionError struct{}

func (e *InterruptionError) Error() string {
	return fmt.Sprint("Interrupted")
}
