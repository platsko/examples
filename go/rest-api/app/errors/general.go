package errors

import (
	"fmt"
)

const (
	ErrCodeUnknown = -1
)

var (
	errNotImplemented = fmt.Errorf("not implemented")
)

func ErrNotImplemented() error {
	return errNotImplemented
}
