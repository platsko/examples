package types

import (
	"database/sql/driver"
)

type (
	// OrgStatus represents status enum column.
	OrgStatus string
)

const (
	OrgStatusActive  OrgStatus = "active"
	OrgStatusBlocked OrgStatus = "blocked"
	OrgStatusUnknown OrgStatus = "unknown"
)

// Scan satisfies scanner interface.
func (t *OrgStatus) Scan(v interface{}) error {
	*t = OrgStatus(v.([]byte))
	return nil
}

// Value satisfies valuer interface.
func (t OrgStatus) Value() (driver.Value, error) {
	return string(t), nil
}
