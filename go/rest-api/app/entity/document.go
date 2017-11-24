package entity

import (
	"github.com/jinzhu/gorm"

	t "lab/go-rest-api/app/repo/types"
)

type (
	// Document describes entity.
	Document struct {
		PrimaryKey                   // use PrimaryKey
		Timestamps                   // use Timestamps
		Organization    Organization // belongs to organization
		OrganizationID  uint64       `gorm:"NOT NULL"` // organization foreign key
		Client          Client       // belongs to client
		ClientID        uint64       `gorm:"NOT NULL"` // client foreign key
		Positions       []Position   // has many positions
		Shift           uint64       `gorm:"NOT NULL"` // shift number
		Number          uint64       `gorm:"NOT NULL"` // document number
		Uid             string       `gorm:"type:VARCHAR(64);UNIQUE_INDEX;NOT NULL"`
		RefUid          string       `gorm:"type:VARCHAR(64);DEFAULT:NULL"`
		Type            t.DocType    `sql:"type:ENUM('receipt','refundReceipt','unknown');DEFAULT:'unknown'"`
		Status          t.DocStatus  `sql:"type:ENUM('confirmed','pending','unknown');DEFAULT:'unknown'"`
		BonusSpent      int64        `gorm:"NOT NULL"`     // amount of spent bonuses
		BonusReceived   int64        `gorm:"NOT NULL"`     // amount of received bonuses
		DiscardedAmount int64        `gorm:"DEFAULT:NULL"` // rounding of amounts
		Payments        []Payment    // has many payments
	}
)

var (
	// Make sure the type satisfies manager interface.
	_ OrmManager = (*Document)(nil)
)

// NewDocument constructs a new entity.
func NewDocument(orm *gorm.DB) *Document {
	e := &Document{}

	// set entity properties and relations
	e.PrimaryKey = primaryKey(orm, e, e.Client, e.Organization)

	return e
}

// TableName returns database table name.
func (*Document) TableName() string {
	return "documents"
}

// Create satisfies entity interface.
func (e *Document) Create() error {
	return e.DB.FirstOrCreate(e).Error
}

// Read satisfies entity interface.
func (e *Document) Read() error {
	return e.DB.First(e).Error
}

// Update satisfies entity interface.
func (e *Document) Update() error {
	return e.DB.Updates(e).Error
}

// Delete satisfies entity interface.
func (e *Document) Delete() error {
	return e.DB.Delete(e).Error
}
