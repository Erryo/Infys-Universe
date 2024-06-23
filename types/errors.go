package types

import "errors"

var (
	ErrValueTooLong    error = errors.New("value exceeds length limit")
	ErrFieldEmpty      error = errors.New("field is empty")
	ErrInvalidArgument error = errors.New("specified argument is invalid")
	ErrAlreadyExists   error = errors.New("row already exists")
	ErrDuplicateKey    error = errors.New("provided key is a duplicate")
	ErrNoRows          error = errors.New("query did not set any rows")
)
