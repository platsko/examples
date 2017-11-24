package types

import (
	json "github.com/json-iterator/go"
)

type (
	// pairLine describes pair text line element.
	pairLine struct {
		Type      string `json:"type"`                // describes the element type
		Left      string `json:"left"`                // text line on the left side
		Right     string `json:"right"`               // text line on the right side
		Separator string `json:"separator,omitempty"` // separator between left and right sides
		Size      string `json:"size,omitempty"`      // text size settings
		Alignment string `json:"alignment,omitempty"` // alignment of both text lines
		Wrapping  string `json:"wrapping,omitempty"`  // the way to wrap text lines
	}
)

// PairLine returns new pair text line.
func PairLine() *pairLine {
	return &pairLine{Type: "pair"}
}

// MarshalJSON satisfies marshaler interface.
func (t *pairLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

// SetLeft sets Left property value.
func (t *pairLine) SetLeft(val string) *pairLine {
	t.Left = val
	return t
}

// SetLeft sets Right property value.
func (t *pairLine) SetRight(val string) *pairLine {
	t.Right = val
	return t
}

// SetSeparator sets Separator property value.
// Value by default " " (a space char).
func (t *pairLine) SetSeparator(val string) *pairLine {
	t.Separator = val
	return t
}

// SetSize sets Size property value.
//
// See "./app/http/respond/types/const.go" for supported values:
// Default - value by default, use hardware settings;
// Small - use font size: 4;
// Medium - use font size: 3;
// Large - use font size: 1.
func (t *pairLine) SetSize(val string) *pairLine {
	switch val {
	case Small, Medium, Large:
		t.Size = val
	default:
		t.Size = ""
	}
	return t
}

// SetAlignment sets Alignment property value.
//
// See "./app/http/respond/types/const.go" for supported values:
// Left - value by default, on the left side;
// Right - on the right side.
func (t *pairLine) SetAlignment(val string) *pairLine {
	switch val {
	case Right:
		t.Alignment = val
	default:
		t.Alignment = ""
	}
	return t
}

// SetWrap sets Wrapping property value.
//
// See "./app/http/respond/types/const.go" for supported values:
// Anywhere - value by default, on any character;
// WordWrap - word-wrap.
func (t *pairLine) SetWrap(val string) *pairLine {
	switch val {
	case WordWrap:
		t.Wrapping = val
	default:
		t.Wrapping = ""
	}
	return t
}
