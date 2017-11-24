package entity

import (
	"github.com/jinzhu/gorm"
)

type (
	// BusinessUnit describes entity.
	BusinessUnit struct {
		PrimaryKey                  // use PrimaryKey
		Timestamps                  // use Timestamps
		Organization   Organization // belongs to organization
		OrganizationID uint64       `gorm:"NOT NULL"`                   // organization foreign key
		Name           string       `gorm:"type:VARCHAR(255);NOT NULL"` // business unit name
	}
)

var (
	// Make sure the type satisfies manager interface.
	_ OrmManager = (*BusinessUnit)(nil)
)

// NewBusinessUnit constructs a new entity.
func NewBusinessUnit(orm *gorm.DB) *BusinessUnit {
	e := &BusinessUnit{}

	// set entity properties and relations
	e.PrimaryKey = primaryKey(orm, e, e.Organization)

	return e
}

// TableName returns database table name.
func (*BusinessUnit) TableName() string {
	return "business_units"
}

// Create satisfies entity interface.
func (e *BusinessUnit) Create() error {
	return e.DB.FirstOrCreate(e).Error
}

// Read satisfies entity interface.
func (e *BusinessUnit) Read() error {
	return e.DB.First(e).Error
}

// Update satisfies entity interface.
func (e *BusinessUnit) Update() error {
	return e.DB.Updates(e).Error
}

// Delete satisfies entity interface.
func (e *BusinessUnit) Delete() error {
	return e.DB.Delete(e).Error
}
