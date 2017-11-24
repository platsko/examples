package types

import (
	json "github.com/json-iterator/go"
)

type (
	// SlipLiner describes a slip line interface.
	SlipLiner interface {
		MarshalJSON() ([]byte, error)
	}

	// Slip is a slice set of slip liner interface.
	Slip []SlipLiner
)

// NewSlip returns a new slip type pointer.
func NewSlip() *Slip {
	sb := make(Slip, 0, 0)
	return &sb
}

// AddLine append a line to the slip set.
func (t *Slip) AddLine(line SlipLiner) *Slip {
	*t = append(*t, line)
	return t
}

// MarshalJSON satisfies marshaler interface.
func (t *Slip) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}
