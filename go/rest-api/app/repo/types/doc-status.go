package types

import (
	"database/sql/driver"
)

type (
	// DocStatus represents document status enum column.
	DocStatus string
)

const (
	DocStatusConfirmed DocStatus = "confirmed"
	DocStatusPending   DocStatus = "pending"
	DocStatusUnknown   DocStatus = "unknown"
)

// Scan satisfies scanner interface.
func (t *DocStatus) Scan(v interface{}) error {
	*t = DocStatus(v.([]byte))
	return nil
}

// Scan satisfies valuer interface.
func (t DocStatus) Value() (driver.Value, error) {
	return string(t), nil
}
