package error

import (
	"fmt"
)

type ServiceConfigLoadError struct {
	Prop string
}

func (k *ServiceConfigLoadError) Error() string {
	return fmt.Sprintf("%s is not defined", k.Prop)
}
