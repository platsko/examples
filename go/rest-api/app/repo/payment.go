package repo

import (
	"lab/go-rest-api/app/entity"
	t "lab/go-rest-api/app/repo/types"
	"lab/go-rest-api/database"
)

type (
	// PaymentRepo wraps entity.
	PaymentRepo struct {
		*entity.Payment
	}
)

// NewClient returns entity repository.
func NewPayment(t t.PayType, m t.PayMode, a t.Amount) *PaymentRepo {
	// create and fill entity
	payment := entity.NewPayment(database.Orm())
	payment.Type = t
	payment.Mode = m
	payment.Amount = a.Int64()

	return &PaymentRepo{Payment: payment}
}
