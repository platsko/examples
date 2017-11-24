package types

import (
	"database/sql/driver"
)

type (
	// PayMode represents payment mode enum column.
	PayMode string
)

const (
	PayModeFiscal    PayMode = "fiscal"
	PayModeNonFiscal PayMode = "nonFiscal"
)

// Scan satisfies scanner interface.
func (t *PayMode) Scan(v interface{}) error {
	*t = PayMode(v.([]byte))
	return nil
}

// Scan satisfies valuer interface.
func (t PayMode) Value() (driver.Value, error) {
	return string(t), nil
}
