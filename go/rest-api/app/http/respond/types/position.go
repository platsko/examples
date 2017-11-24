package types

import (
	json "github.com/json-iterator/go"
)

type (
	// position describes response element.
	position struct {
		Index          uint64  `json:"index"`
		DiscountAmount float64 `json:"discountAmount"`
	}
)

// Position returns position element.
func Position(index uint64, discount float64) *position {
	return &position{
		Index:          index,
		DiscountAmount: discount,
	}
}

// MarshalJSON satisfies marshaler interface.
func (t *position) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}
