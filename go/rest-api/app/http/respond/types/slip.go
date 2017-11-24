package types

import (
	json "github.com/json-iterator/go"
)

type (
	// Slipper describes a slip line interface.
	Slipper interface {
		MarshalJSON() ([]byte, error)
	}

	// Slip a slice set of slipper interface
	// to display or printing information.
	Slip []Slipper
)

// NewSlip returns a new slip type pointer.
func NewSlip() *Slip {
	sb := make(Slip, 0, 0)
	return &sb
}

// AddLine append a line to the slip set.
func (sb *Slip) AddLine(line Slipper) *Slip {
	*sb = append(*sb, line)
	return sb
}

// MarshalJSON satisfies marshaler interface.
func (sb *Slip) MarshalJSON() ([]byte, error) {
	return json.Marshal(*sb)
}
