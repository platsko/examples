// Copyright © 2020 The EVEN Lab Team

package types

import (
	"strings"
)

type (
	// PhoneNumber represents a client phone number.
	PhoneNumber string
)

// @FIXME there is a stub.
// @TODO implementation instead.
// Parse tries to parse a string to PhoneNumber type
func (n *PhoneNumber) Parse(v string) (*PhoneNumber, error) {
	*n = PhoneNumber(v)
	return n, nil
}

// String satisfies Stringer interface.
func (n *PhoneNumber) String() string {
	return string(*n)
}

// Value returns cleared value as a string.
func (n *PhoneNumber) Value() string {
	return strings.TrimLeft(n.String(), "+")
}
