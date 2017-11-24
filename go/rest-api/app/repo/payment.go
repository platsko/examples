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

// NewPayment returns entity repository.
func NewPayment(payType t.PayType, payMode t.PayMode, amount t.Amount) *PaymentRepo {
	// create and fill entity
	payment := entity.NewPayment(database.Orm())
	payment.Type = payType
	payment.Mode = payMode
	payment.Amount = amount.Int64()

	return &PaymentRepo{Payment: payment}
}
