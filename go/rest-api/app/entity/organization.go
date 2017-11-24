package entity

import (
	"github.com/jinzhu/gorm"

	t "lab/go-rest-api/app/repo/types"
)

type (
	// Organization describes entity.
	Organization struct {
		PrimaryKey             // use PrimaryKey
		Timestamps             // use Timestamps
		Name       string      `gorm:"type:VARCHAR(255);NOT NULL"` // organization name
		Token      string      `gorm:"type:VARCHAR(64);NOT NULL"` // organization access token
		Status     t.OrgStatus `sql:"type:ENUM('active','blocked','unknown');DEFAULT:'unknown'"` // organization status
	}
)

var (
	// Make sure the type satisfies manager interface.
	_ OrmManager = (*Organization)(nil)
)

// NewOrganization constructs a new entity.
func NewOrganization(orm *gorm.DB) *Organization {
	e := &Organization{}

	// set entity properties and relations
	e.PrimaryKey = primaryKey(orm, e)

	return e
}

// TableName returns database table name.
func (*Organization) TableName() string {
	return "organizations"
}

// Create satisfies entity interface.
func (e *Organization) Create() error {
	return e.DB.FirstOrCreate(e).Error
}

// Read satisfies entity interface.
func (e *Organization) Read() error {
	return e.DB.First(e).Error
}

// Update satisfies entity interface.
func (e *Organization) Update() error {
	return e.DB.Updates(e).Error
}

// Delete satisfies entity interface.
func (e *Organization) Delete() error {
	return e.DB.Delete(e).Error
}
