package types

import (
	json "github.com/json-iterator/go"
)

type (
	// textLine describes text line element.
	textLine struct {
		Type      string `json:"type"`                // describes the element type
		Text      string `json:"text"`                // text line
		Size      string `json:"size,omitempty"`      // text size
		Alignment string `json:"alignment,omitempty"` // alignment of text
		Wrapping  string `json:"wrapping,omitempty"`  // the way to wrap text
	}
)

// TextLine returns a new text line.
func TextLine() *textLine {
	return &textLine{Type: "text"}
}

// MarshalJSON satisfies marshaler interface.
func (t *textLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

// SetText sets Text property value.
func (t *textLine) SetText(val string) *textLine {
	t.Text = val
	return t
}

// SetSize sets Size property value.
//
// See "./app/http/respond/types/const.go" for supported values:
// Default - value by default, use hardware settings;
// Small - use font size: 4;
// Medium - use font size: 3;
// Large - use font size: 1.
func (t *textLine) SetSize(val string) *textLine {
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
func (t *textLine) SetAlignment(val string) *textLine {
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
func (t *textLine) SetWrap(val string) *textLine {
	switch val {
	case WordWrap:
		t.Wrapping = val
	default:
		t.Wrapping = ""
	}
	return t
}
