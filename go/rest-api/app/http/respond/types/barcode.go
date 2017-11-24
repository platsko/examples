package types

import (
	json "github.com/json-iterator/go"
)

type (
	// barCode describes bar-code line element.
	barCode struct {
		Type      string `json:"type"`                // describes the element type
		Data      string `json:"data"`                // data for encoding in EAN13
		Size      string `json:"size,omitempty"`      // bar-code size
		PrintText string `json:"printText,omitempty"` // print string value of barcode on the form
	}
)

// BarCode returns a new bar-code line.
// Valid data are 12 digits without check number,
// or 13 digits with check number.
func BarCode(data string) *barCode {
	return &barCode{Type: "barcode", Data: data}
}

// MarshalJSON satisfies marshaler interface.
func (t *barCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

// SetSize sets Size property value.
//
// See "./app/http/respond/types/const.go" for supported values:
// Tiny;
// Small;
// Normal - value by default;
// Large;
// Extralarge.
func (t *barCode) SetSize(val string) *barCode {
	switch val {
	case Tiny, Small, Large, Extra:
		t.Size = val
	default:
		t.Size = ""
	}
	return t
}

// SetPrintText sets PrintText property value.
// The property depends on specific hardware model.
//
// See "./app/http/respond/types/const.go" for supported values:
// None - value by default, do not print;
// Above - above the barcode;
// Below - below the barcode;
// Everywhere - everywhere on form.
func (t *barCode) SetPrintText(val string) *barCode {
	switch val {
	case Above, Below, Everywhere:
		t.PrintText = val
	default:
		t.PrintText = ""
	}
	return t
}
