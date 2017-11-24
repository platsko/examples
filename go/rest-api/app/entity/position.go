package entity

import (
	"github.com/jinzhu/gorm"
)

type (
	// Position describes entity.
	Position struct {
		PrimaryKey              // use PrimaryKey
		Timestamps              // use Timestamps
		Document       Document // belongs to document
		DocumentID     uint64   `gorm:"NOT NULL"`                   // document foreign key
		Index          uint64   `gorm:"NOT NULL"`                   // position index in document
		Uid            string   `gorm:"type:VARCHAR(64);NOT NULL"`  // position uid
		Text           string   `gorm:"type:VARCHAR(255);NOT NULL"` // description
		Price          int64    `gorm:"NOT NULL"`                   // price amount
		MinimumPrice   int64    `gorm:"NOT NULL"`                   // minimum price amount
		Quantity       string   `gorm:"type:VARCHAR(11);NOT NULL"`  // quantity of position
		TotalAmount    int64    `gorm:"NOT NULL"`                   // total amount
		DiscountAmount int64    `gorm:"NOT NULL"`                   // amount of discount
		PaidAmount     int64    `gorm:"NOT NULL"`                   // amount of paid
		BonusReceived  int64    `gorm:"NOT NULL"`                   // amount of received bonuses
	}
)

var (
	// Make sure the type satisfies manager interface.
	_ OrmManager = (*Position)(nil)
)

// NewPosition constructs a new entity.
func NewPosition(orm *gorm.DB) *Position {
	e := &Position{}

	// set entity properties and relations
	e.PrimaryKey = primaryKey(orm, e, e.Document)

	return e
}

// TableName returns database table name.
func (*Position) TableName() string {
	return "positions"
}

// Create satisfies entity interface.
func (e *Position) Create() error {
	return e.DB.FirstOrCreate(e).Error
}

// Read satisfies entity interface.
func (e *Position) Read() error {
	return e.DB.First(e).Error
}

// Update satisfies entity interface.
func (e *Position) Update() error {
	return e.DB.Updates(e).Error
}

// Delete satisfies entity interface.
func (e *Position) Delete() error {
	return e.DB.Delete(e).Error
}
