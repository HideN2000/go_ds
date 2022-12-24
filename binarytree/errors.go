package binarytree

import "errors"

var (
	ErrInvalidIndex = errors.New("index is not valid")
	ErrTreeEmpty    = errors.New("tree is empty")
	ErrNotFound     = errors.New("value is not found")
	ErrUnexpected   = errors.New("unexpected error occured. A bug may be in the source code")
)
