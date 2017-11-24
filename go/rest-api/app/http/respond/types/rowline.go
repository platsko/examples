package types

import (
	json "github.com/json-iterator/go"
)

type (
	// rowLine describes row line element.
	rowLine struct {
		// Type describes the element type.
		Type string `json:"type"`

		// The symbol to be used as a separator.
		// The property depends on specific hardware model.
		// Value by default "-".
		Symbol string `json:"symbol"`

		// Size of separator settings.
		// The property depends on specific hardware model.
		Size string `json:"size,omitempty"`
	}
)

// RowLine returns a new row line.
func RowLine() *rowLine {
	return &rowLine{Type: "line", Symbol: "-"}
}

// MarshalJSON satisfies marshaler interface.
func (r *rowLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}

// SetSymbol sets Symbol property value.
func (r *rowLine) SetSymbol(val string) *rowLine {
	r.Symbol = val
	return r
}

// SetSize sets Size property value.
// @SEE "./app/http/respond/types/const.go" for supported values:
// Default - value by default, use hardware settings;
// Small - use font size: 4;
// Medium - use font size: 3;
// Large - use font size: 1.
func (r *rowLine) SetSize(val string) *rowLine {
	switch val {
	case Small, Medium, Large:
		r.Size = val
	default:
		r.Size = ""
	}
	return r
}
