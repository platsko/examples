package entity

import (
	"github.com/jinzhu/gorm"
)

type (
	// WorkPlace describes entity.
	WorkPlace struct {
		PrimaryKey                  // use PrimaryKey
		Timestamps                  // use Timestamps
		BusinessUnit   BusinessUnit // belongs to business unit
		BusinessUnitID uint64       `gorm:"NOT NULL"`                   // business unit foreign key
		Name           string       `gorm:"type:VARCHAR(255);NOT NULL"` // work place name
	}
)

var (
	// Make sure the type satisfies manager interface.
	_ OrmManager = (*WorkPlace)(nil)
)

// WorkPlace constructs a new entity.
func NewWorkPlace(orm *gorm.DB) *WorkPlace {
	e := &WorkPlace{}

	// set entity properties and relations
	e.PrimaryKey = primaryKey(orm, e, e.BusinessUnit)

	return e
}

// TableName returns database table name.
func (*WorkPlace) TableName() string {
	return "work_places"
}

// Create satisfies entity interface.
func (e *WorkPlace) Create() error {
	return e.DB.FirstOrCreate(e).Error
}

// Read satisfies entity interface.
func (e *WorkPlace) Read() error {
	return e.DB.First(e).Error
}

// Update satisfies entity interface.
func (e *WorkPlace) Update() error {
	return e.DB.Updates(e).Error
}

// Delete satisfies entity interface.
func (e *WorkPlace) Delete() error {
	return e.DB.Delete(e).Error
}
