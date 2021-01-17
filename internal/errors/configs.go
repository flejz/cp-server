package errors

import (
	"fmt"
)

type ServiceConfigLoadError struct {
	Prop string
}

func (c *ServiceConfigLoadError) Error() string {
	return fmt.Sprintf("%s is not defined", c.Prop)
}

// ---

type ServerConfigLoadError struct {
	Prop string
}

func (c *ServerConfigLoadError) Error() string {
	return fmt.Sprintf("%s is not defined", c.Prop)
}
