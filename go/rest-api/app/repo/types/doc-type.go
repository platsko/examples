package types

import (
	"database/sql/driver"
)

type (
	// DocType represents document type enum column.
	DocType string
)

const (
	DocTypeReceipt       DocType = "receipt"
	DocTypeRefundReceipt DocType = "refundReceipt"
	DocTypeUnknown       DocType = "unknown"
)

// Scan satisfies scanner interface.
func (t *DocType) Scan(v interface{}) error {
	*t = DocType(v.([]byte))
	return nil
}

// Scan satisfies valuer interface.
func (t DocType) Value() (driver.Value, error) {
	return string(t), nil
}
