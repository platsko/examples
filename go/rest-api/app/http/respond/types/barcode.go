package types

import (
	json "github.com/json-iterator/go"
)

type (
	// barCode describes bar-code line element.
	barCode struct {
		// Type describes the element type.
		Type string `json:"type"`

		// Data for encoding in EAN13.
		// Valid data are 12 digits without check number,
		// or 13 digits with check number.
		Data string `json:"data"`

		// Bar-Code size.
		Size string `json:"size,omitempty"`

		// Print a string value of barcode on the form.
		// The property depends on specific hardware model.
		PrintText string `json:"printText,omitempty"`
	}
)

// BarCode returns a new bar-code line.
func BarCode(data string) *barCode {
	return &barCode{Type: "barcode", Data: data}
}

// MarshalJSON satisfies marshaler interface.
func (c *barCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}

// SetSize sets Size property value.
// @SEE "./app/http/respond/types/const.go" for supported values:
// Tiny;
// Small;
// Normal - value by default;
// Large;
// Extralarge.
func (c *barCode) SetSize(val string) *barCode {
	switch val {
	case Tiny, Small, Large, Extra:
		c.Size = val
	default:
		c.Size = ""
	}
	return c
}

// SetPrintText sets PrintText property value.
// @SEE "./app/http/respond/types/const.go" for supported values:
// None - value by default, do not print;
// Above - above the barcode;
// Below - below the barcode;
// Everywhere - everywhere on form.
func (c *barCode) SetPrintText(val string) *barCode {
	switch val {
	case Above, Below, Everywhere:
		c.PrintText = val
	default:
		c.PrintText = ""
	}
	return c
}
