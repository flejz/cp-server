package error

import (
	"fmt"
)

type InvalidCredentialsError struct{}

func (k *InvalidCredentialsError) Error() string {
	return fmt.Sprint("Credentials are invalid")
}

type UserExistsError struct{}

func (k *UserExistsError) Error() string {
	return fmt.Sprint("User already exists")
}
