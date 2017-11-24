package errors

import (
	"fmt"
)

const (
	ErrCodeUnknown = -1
)

var (
	errNotImplemented = fmt.Errorf("not implemented")
	errNotEnoughFunds = fmt.Errorf("not enough funds")
	errNotSupported   = fmt.Errorf("not supported")
)

func ErrNotImplemented() error {
	return errNotImplemented
}

func ErrNotEnoughFunds() error {
	return errNotEnoughFunds
}

func ErrNotSupported() error {
	return errNotSupported
}
