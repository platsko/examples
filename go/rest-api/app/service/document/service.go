package document

import (
	"lab/go-rest-api/app/entity"
	"lab/go-rest-api/app/errors"
	"lab/go-rest-api/app/http/request"
	"lab/go-rest-api/app/repo"
	"lab/go-rest-api/app/repo/types"
)

// Confirm makes the confirm action.
func Confirm(rp *repo.DocumentRepo, f *request.DocumentForm) error {
	if len(f.Payments) == 0 {
		return errors.ErrNotSupported()
	}

	// append payments to document
	for _, p := range f.Payments {
		if p.Mode == types.PayModeNonFiscal && p.Amount.Int64() != rp.BonusSpent {
			return errors.ErrNotSupported()
		}
		rp.Document.Payments = append(
			rp.Document.Payments,
			*repo.NewPayment(p.Type, p.Mode, p.Amount).Payment,
		)
	}

	return nil
}

// PayByBonus makes the pay-by-bonus action.
func PayByBonus(rp *repo.DocumentRepo, f *request.DocumentForm) error {
	if f.Payment.Mode != types.PayModeNonFiscal {
		return errors.ErrNotSupported()
	}

	// check client bonus account
	bonus := &rp.Client.ClientBonus
	amount := f.Payment.Amount.Int64()
	if amount > bonus.Amount {
		return errors.ErrNotEnoughFunds()
	}

	// makes distribution amount
	if err := distributeBonusAmount(rp.Positions, amount); err != nil {
		return err
	}

	// debit bonus amount
	rp.BonusSpent = amount
	bonus.Amount -= amount

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
