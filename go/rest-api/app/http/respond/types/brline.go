package types

import (
	json "github.com/json-iterator/go"
)

type (
	// brLine describes br line element type.
	brLine struct {
		Type string `json:"type"` // describes the element type

	}
)

// Br returns a new br line.
func Br() *brLine {
	return &brLine{Type: "br"}
}

// MarshalJSON satisfies marshaler interface.
func (t *brLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}
