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
func (i *Print) AddSlip(s *Slip) *Print {
	*i = append(*i, s)
	return i
}

// MarshalJSON satisfies marshaler interface.
func (i *Print) MarshalJSON() ([]byte, error) {
	return json.Marshal(*i)
}
