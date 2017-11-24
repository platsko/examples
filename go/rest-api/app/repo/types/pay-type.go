package types

import (
	"database/sql/driver"
)

type (
	// PayType represents payment type enum column.
	PayType string
)

const (
	PayTypeBonus PayType = "bonus"
	PayTypeCash  PayType = "cash"
	PayTypeCard  PayType = "creditCard"
	PayTypeOther PayType = "other"
)

// Scan satisfies scanner interface.
func (t *PayType) Scan(v interface{}) error {
	*t = PayType(v.([]byte))
	return nil
}

// Scan satisfies valuer interface.
func (t PayType) Value() (driver.Value, error) {
	return string(t), nil
}
