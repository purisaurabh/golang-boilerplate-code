package errors

import "errors"

var (
	ErrParameterMissing = errors.New("parameter missing")
	ErrInvalidFormat    = errors.New("invalid request format")
)
