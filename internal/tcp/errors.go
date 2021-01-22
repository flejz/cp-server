package tcp

import (
	"errors"
)

var ErrInvalidPort = errors.New("invalid port")
var ErrInvalid = errors.New("invalid")
var ErrInterrupted = errors.New("interrupted")
