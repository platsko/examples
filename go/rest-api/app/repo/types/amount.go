package types

import (
	"fmt"
	"math/big"
)

type (
	// Amount represents money amount.
	Amount int64
)

var (
	// scaleToUnits represents scale amount to units.
	scaleToUnits = big.NewFloat(100)
)

// NewAmount returns Amount for specified value.
func NewAmount(amount int64) *Amount {
	a := Amount(amount)
	return &a
}

// Float64 returns Amount converted to float64.
func (t *Amount) Float64() float64 {
	f := big.NewFloat(float64(*t))
	a, _ := f.Quo(f, scaleToUnits).Float64()
	return a
}

// Int64 returns Amount scaled and converted to int64.
func (t *Amount) Int64() int64 {
	return int64(*t)
}

// Parse tries to parse a string to Amount.
func (t *Amount) Parse(v string) (*Amount, error) {
	f, _, err := big.ParseFloat(v, 10, 2, big.AwayFromZero)
	if err != nil {
		return nil, fmt.Errorf("parse float: %w", err)
	}

	amount, accuracy := f.Mul(f, scaleToUnits).Int64()
	if accuracy != 0 {
		return nil, fmt.Errorf("not accuracy")
	}

	*t = Amount(amount)

	return t, nil
}

// String satisfies stringer interface.
func (t *Amount) String() string {
	f := big.NewFloat(float64(*t))
	return f.Quo(f, scaleToUnits).Text('f', 2)
}
