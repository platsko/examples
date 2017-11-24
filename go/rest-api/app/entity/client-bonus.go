package entity

import (
	"github.com/jinzhu/gorm"
)

type (
	// ClientBonus describes entity.
	ClientBonus struct {
		PrimaryKey        // use PrimaryKey
		Timestamps        // use Timestamps
		ClientID   uint64 `gorm:"NOT NULL"` // client foreign key
		Amount     int64  `gorm:"NOT NULL"` // amount of received bonuses
	}
)

var (
	// Make sure the type satisfies manager interface.
	_ OrmManager = (*ClientBonus)(nil)
)

// NewClientBonus constructs a new entity.
func NewClientBonus(orm *gorm.DB) *ClientBonus {
	e := &ClientBonus{}

	// set entity properties and relations
	e.PrimaryKey = primaryKey(orm, e)

	return e
}

// TableName returns database table name.
func (*ClientBonus) TableName() string {
	return "client_bonuses"
}

// Create satisfies entity interface.
func (e *ClientBonus) Create() error {
	return e.DB.FirstOrCreate(e).Error
}

// Read satisfies entity interface.
func (e *ClientBonus) Read() error {
	return e.DB.First(e).Error
}

// Update satisfies entity interface.
func (e *ClientBonus) Update() error {
	return e.DB.Updates(e).Error
}

// Delete satisfies entity interface.
func (e *ClientBonus) Delete() error {
	return e.DB.Delete(e).Error
}
