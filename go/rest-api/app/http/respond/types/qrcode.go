package types

import (
	json "github.com/json-iterator/go"
)

type (
	// qrCode describes qr-code line element.
	qrCode struct {
		Type       string `json:"type"`                 // describes the element type
		Data       string `json:"data"`                 // data for encoding in a qr-code
		Size       string `json:"size,omitempty"`       // qr-code size
		Correction string `json:"correction,omitempty"` // degree of correction (noise immunity)
	}
)

// QrCode returns a new qr-code line.
func QrCode(data string) *qrCode {
	return &qrCode{Type: "qrcode", Data: data}
}

// MarshalJSON satisfies marshaler interface.
func (t *qrCode) MarshalJSON() ([]byte, error) {
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
func (t *qrCode) SetSize(val string) *qrCode {
	switch val {
	case Tiny, Small, Large, Extra:
		t.Size = val
	default:
		t.Size = ""
	}
	return t
}

// SetCorrection sets Correction property value.
// The property depends on specific hardware model.
// For small size, do not specify ultra / high values
// as it may result in a QR code recognition error.
//
// See "./app/http/respond/types/const.go" for supported values:
// Low;
// Medium - value by default;
// High;
// Ultra.
func (t *qrCode) SetCorrection(val string) *qrCode {
	switch val {
	case Low, High, Ultra:
		t.Correction = val
	default:
		t.Correction = ""
	}
	return t
}
