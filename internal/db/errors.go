package db

import (
	"errors"
)

var ErrInvalidMemoryName = errors.New("invalid memory name")
var ErrInvalidSQLitePath = errors.New("invalid sqlite path")
var ErrInvalidType = errors.New("invalid db type")
