package types

import (
	json "github.com/json-iterator/go"
)

type (
	// textLine describes text line element.
	textLine struct {
		// Type describes the element type.
		Type string `json:"type"`

		// The text line to display or printing.
		// The property depends on specific hardware model.
		Text string `json:"text"`

		// Text size settings.
		// The property depends on specific hardware model.
		Size string `json:"size,omitempty"`

		// Alignment of text lines.
		Alignment string `json:"alignment,omitempty"`

		// The way to wrap text lines.
		Wrapping string `json:"wrapping,omitempty"`
	}
)

// TextLine returns a new text line.
func TextLine() *textLine {
	return &textLine{Type: "text"}
}

// MarshalJSON satisfies marshaler interface.
func (l *textLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(*l)
}

// SetText sets Text property value.
func (l *textLine) SetText(val string) *textLine {
	l.Text = val
	return l
}

// SetSize sets Size property value.
// @SEE "./app/http/respond/types/const.go" for supported values:
// Default - value by default, use hardware settings;
// Small - use font size: 4;
// Medium - use font size: 3;
// Large - use font size: 1.
func (l *textLine) SetSize(val string) *textLine {
	switch val {
	case Small, Medium, Large:
		l.Size = val
	default:
		l.Size = ""
	}
	return l
}

// SetAlignment sets Alignment property value.
// @SEE "./app/http/respond/types/const.go" for supported values:
// Left - value by default, on the left side;
// Right - on the right side.
func (l *textLine) SetAlignment(val string) *textLine {
	switch val {
	case Right:
		l.Alignment = val
	default:
		l.Alignment = ""
	}
	return l
}

// SetWrap sets Wrapping property value.
// @SEE "./app/http/respond/types/const.go" for supported values:
// Anywhere - value by default, on any character;
// WordWrap - word-wrap.
func (l *textLine) SetWrap(val string) *textLine {
	switch val {
	case WordWrap:
		l.Wrapping = val
	default:
		l.Wrapping = ""
	}
	return l
}
