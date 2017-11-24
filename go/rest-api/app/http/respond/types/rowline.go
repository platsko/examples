package types

import (
	json "github.com/json-iterator/go"
)

type (
	// rowLine describes row line element.
	rowLine struct {
		Type   string `json:"type"`           // describes the element type
		Symbol string `json:"symbol"`         // symbol to be used as a separator
		Size   string `json:"size,omitempty"` // size of separator settings
	}
)

// RowLine returns a new row line.
func RowLine() *rowLine {
	return &rowLine{Type: "line", Symbol: "-"}
}

// MarshalJSON satisfies marshaler interface.
func (t *rowLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

// SetSymbol sets Symbol property value.
// Value by default "-".
func (t *rowLine) SetSymbol(val string) *rowLine {
	t.Symbol = val
	return t
}

// SetSize sets Size property value.
//
// See "./app/http/respond/types/const.go" for supported values:
// Default - value by default, use hardware settings;
// Small - use font size: 4;
// Medium - use font size: 3;
// Large - use font size: 1.
func (t *rowLine) SetSize(val string) *rowLine {
	switch val {
	case Small, Medium, Large:
		t.Size = val
	default:
		t.Size = ""
	}
	return t
}
