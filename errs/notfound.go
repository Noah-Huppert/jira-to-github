package errs

import (
	"errors"
)

// ErrNotFound indicates the API model being requested was not found
var ErrNotFound error = errors.New("not found")
