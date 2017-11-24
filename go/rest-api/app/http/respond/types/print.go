package types

import (
	json "github.com/json-iterator/go"
)

type (
	// Print a slice set of slip type.
	Print []*Slip
)

// NewPrint returns a new print set.
func NewPrint() *Print {
	d := make(Print, 0, 0)
	return &d
}

// AddSlip appends a slip to print set.
func (t *Print) AddSlip(s *Slip) *Print {
	*t = append(*t, s)
	return t
}

// MarshalJSON satisfies marshaler interface.
func (t *Print) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}
