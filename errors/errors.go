package lib

import "errors"

var (
	ErrInvalidIndex = errors.New("index is not valid")
	ErrInvalidValue = errors.New("value is not valid")
	ErrNotFound     = errors.New("value is not found")
	ErrUnexpected   = errors.New("unexpected error occured. A bug may be in the source code")
)
