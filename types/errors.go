package types

import "errors"

var (
	ErrValueTooLong    error = errors.New("infys:value exceeds length limit")
	ErrFieldEmpty      error = errors.New("infys:field is empty")
	ErrInvalidArgument error = errors.New("infys:specified argument is invalid")
	ErrAlreadyExists   error = errors.New("infys:row already exists")
	ErrDuplicateKey    error = errors.New("infys:provided key is a duplicate")
	ErrNoRows          error = errors.New("infys:query did not set any rows")
	ErrFKViolated      error = errors.New("infys:foreign key constraint violated")
)
