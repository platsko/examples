package request

import (
	t "lab/go-rest-api/app/repo/types"
)

type (
	// Payment describes payment form field.
	Payment struct {
		Type   t.PayType `json:"type"`                       // payment type
		Mode   t.PayMode `json:"mode" validate:"required"`   // payment mode
		Amount t.Amount  `json:"amount" validate:"required"` // payment amount
	}
)
