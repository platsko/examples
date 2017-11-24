package types

import (
	json "github.com/json-iterator/go"
)

type (
	// pairLine describes pair text line element.
	pairLine struct {
		// Type describes the element type.
		Type string `json:"type"`

		// The text line to display or printing on the left side.
		// The property depends on specific hardware model.
		Left string `json:"left"`

		// The text line to display or printing on the right side.
		// The property depends on specific hardware model.
		Right string `json:"right"`

		// The separator between the text
		// on the right and left sides.
		// The property depends on specific hardware model.
		// Value by default " " (a space char).
		Separator string `json:"separator,omitempty"`

		// Text size settings.
		// The property depends on specific hardware model.
		Size string `json:"size,omitempty"`

		// Alignment of both text lines.
		Alignment string `json:"alignment,omitempty"`

		// The way to wrap text lines.
		Wrapping string `json:"wrapping,omitempty"`
	}
)

// PairLine returns new pair text line.
func PairLine() *pairLine {
	return &pairLine{Type: "pair"}
}

// MarshalJSON satisfies marshaler interface.
func (l *pairLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(*l)
}

// SetLeft sets Left property value.
func (l *pairLine) SetLeft(val string) *pairLine {
	l.Left = val
	return l
}

// SetLeft sets Right property value.
func (l *pairLine) SetRight(val string) *pairLine {
	l.Right = val
	return l
}

// SetSeparator sets Separator property value.
func (l *pairLine) SetSeparator(val string) *pairLine {
	l.Separator = val
	return l
}

// SetSize sets Size property value.
// @SEE "./app/http/respond/types/const.go" for supported values:
// Default - value by default, use hardware settings;
// Small - use font size: 4;
// Medium - use font size: 3;
// Large - use font size: 1.
func (l *pairLine) SetSize(val string) *pairLine {
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
func (l *pairLine) SetAlignment(val string) *pairLine {
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
func (l *pairLine) SetWrap(val string) *pairLine {
	switch val {
	case WordWrap:
		l.Wrapping = val
	default:
		l.Wrapping = ""
	}
	return l
}
