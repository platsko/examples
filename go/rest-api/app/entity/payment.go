package entity

import (
	"github.com/jinzhu/gorm"

	"lab/go-rest-api/app/errors"
	t "lab/go-rest-api/app/repo/types"
)

type (
	// Payment describes entity.
	Payment struct {
		PrimaryKey           // use PrimaryKey
		Timestamps           // use Timestamps
		Document   Document  // belongs to document
		DocumentID uint64    `gorm:"NOT NULL"` // document foreign key
		Type       t.PayType `sql:"type:ENUM('bonus','cash','creditCard','other');DEFAULT:'other'"`
		Mode       t.PayMode `sql:"type:ENUM('fiscal','nonFiscal','unknown');DEFAULT:'unknown'"`
		Amount     int64     `gorm:"NOT NULL"` // payment amount
	}
)

var (
	// Make sure the type satisfies manager interface.
	_ OrmManager = (*Payment)(nil)
)

// NewPayment constructs a new entity.
func NewPayment(orm *gorm.DB) *Payment {
	e := &Payment{}

	// set entity properties and relations
	e.PrimaryKey = primaryKey(orm, e, e.Document)

	return e
}

// TableName returns database table name.
func (*Payment) TableName() string {
	return "payments"
}

// Create satisfies entity interface.
func (e *Payment) Create() error {
	return e.DB.FirstOrCreate(e).Error
}

// Read satisfies entity interface.
func (e *Payment) Read() error {
	return e.DB.First(e).Error
}

// Update satisfies entity interface.
func (e *Payment) Update() error {
	return errors.ErrNotSupported()
}

// Delete satisfies entity interface.
func (e *Payment) Delete() error {
	return errors.ErrNotSupported()
}
