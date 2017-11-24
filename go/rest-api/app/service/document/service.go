package document

import (
	"lab/go-rest-api/app/entity"
	"lab/go-rest-api/app/errors"
	"lab/go-rest-api/app/http/request"
	"lab/go-rest-api/app/repo"
)

// Calculate makes the calculate action.
func Calculate(rp *repo.DocumentRepo) error {
	// make distribution amount
	if err := distributeBonusAmount(
		rp.Positions,
		rp.Client.ClientBonus.Amount,
	); err != nil {
		return err
	}

	return nil
}

// PayByBonus makes the pay-by-bonus action.
func PayByBonus(rp *repo.DocumentRepo, f *request.DocumentForm) error {
	// check client bonus account
	amount := f.Payment.Amount.Int64()
	if amount > rp.Client.ClientBonus.Amount {
		return errors.ErrNotEnoughFunds()
	}

	// debit bonus amount
	rp.BonusSpent = amount
	rp.Client.ClientBonus.Amount -= amount

	return nil
}

// CancelBonusPayment makes the cancel-bonus-payment action.
func CancelBonusPayment(rp *repo.DocumentRepo, f *request.DocumentForm) error {
	// credit bonus amount
	amount := f.Payment.Amount.Int64()
	rp.BonusReceived = amount
	rp.Client.ClientBonus.Amount += amount

	return nil
}

// Confirm makes the confirm action.
func Confirm(rp *repo.DocumentRepo, f *request.DocumentForm) error {
	// append payments to document
	for _, p := range f.Payments {
		rp.Document.Payments = append(
			rp.Document.Payments,
			*repo.NewPayment(p.Type, p.Mode, p.Amount).Payment,
		)
	}

	return nil
}

// @TODO make sure that this method performs the correct distribution.
// distributeBonusAmount makes distribution amount by document positions.
func distributeBonusAmount(positions []entity.Position, amount int64) error {
	// iterate by positions list
	for _, p := range positions {
		maxPosBonus := p.Price - p.MinimumPrice
		if maxPosBonus > 0 && maxPosBonus <= amount {
			amount -= maxPosBonus
			p.DiscountAmount = maxPosBonus
		}
		if amount == 0 {
			break
		}
	}

	return nil
}
