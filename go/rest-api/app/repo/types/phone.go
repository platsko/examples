package types

import (
	"fmt"
	"strings"
)

type (
	// Phone represents phone number.
	Phone string
)

// @FIXME there is a stub.
// @TODO implementation instead.
// Parse tries to parse a string to Phone.
func (t *Phone) Parse(v string) (*Phone, error) {
	if len(v) != 12 || !strings.HasPrefix(v, "+") {
		return nil, fmt.Errorf("invalid format")
	}

	*t = Phone(v)

	return t, nil
}

// String satisfies stringer interface.
func (t *Phone) String() string {
	return string(*t)
}

// Value returns cleared value as a string.
func (t *Phone) Value() string {
	return strings.TrimLeft(t.String(), "+")
}
