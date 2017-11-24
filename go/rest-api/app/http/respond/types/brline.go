package types

import (
	json "github.com/json-iterator/go"
)

type (
	// brLine describes br line element type.
	brLine struct {
		// Type describes the element type.
		Type string `json:"type"`
	}
)

// BR returns a new br line.
func BR() *brLine {
	return &brLine{Type: "br"}
}

// MarshalJSON satisfies marshaler interface.
func (br *brLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(*br)
}
