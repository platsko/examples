package types

import (
	json "github.com/json-iterator/go"
)

type (
	// qrCode describes qr-code line element.
	qrCode struct {
		// Type describes the element type.
		Type string `json:"type"`

		// Data for encoding in a qr-code.
		Data string `json:"data"`

		// qr-code size.
		Size string `json:"size,omitempty"`

		// The degree of correction (noise immunity).
		// The property depends on specific hardware model.
		// For small size, do not specify ultra / high values
		// as it may result in a QR code recognition error.
		Correction string `json:"correction,omitempty"`
	}
)

// QrCode returns a new qr-code line.
func QrCode(data string) *qrCode {
	return &qrCode{Type: "qrcode", Data: data}
}

// MarshalJSON satisfies marshaler interface.
func (c *qrCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}

// SetSize sets Size property value.
// @SEE "./app/http/respond/types/const.go" for supported values:
// Tiny;
// Small;
// Normal - value by default;
// Large;
// Extralarge.
func (c *qrCode) SetSize(val string) *qrCode {
	switch val {
	case Tiny, Small, Large, Extra:
		c.Size = val
	default:
		c.Size = ""
	}
	return c
}

// SetCorrection sets Correction property value.
// @SEE "./app/http/respond/types/const.go" for supported values:
// Low;
// Medium - value by default;
// High;
// Ultra.
func (c *qrCode) SetCorrection(val string) *qrCode {
	switch val {
	case Low, High, Ultra:
		c.Correction = val
	default:
		c.Correction = ""
	}
	return c
}
