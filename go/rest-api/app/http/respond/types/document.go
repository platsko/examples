package types

import (
	json "github.com/json-iterator/go"

	"lab/go-rest-api/app/repo"
	rt "lab/go-rest-api/app/repo/types"
)

type (
	// Document describes response element.
	Document struct {
		Positions []position `json:"positions"`
	}
)

// Position returns position element.
func NewDocument(rp *repo.DocumentRepo) *Document {
	e := &Document{Positions: make([]position, 0)}
	for _, p := range rp.Positions {
		discount := rt.NewAmount(p.DiscountAmount).Float64()
		e.Positions = append(e.Positions, *Position(p.Index, discount))
	}
	return e
}

// MarshalJSON satisfies marshaler interface.
func (t *Document) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}
