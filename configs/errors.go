package config

import (
	"errors"
	"fmt"
)

var ErrInvalidDBType = errors.New("invalid db type")
var ErrInvalidConfig = errors.New("invalid config")

type ErrEnvNotFound struct{ envName }

func (e *ErrEnvNotFound) Error() string { return fmt.Sprintf("env not found: %s", e.envName) }
