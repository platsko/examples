package entity

import (
	"github.com/jinzhu/gorm"
)

type (
	// Client describes entity.
	Client struct {
		PrimaryKey         // use PrimaryKey
		Timestamps         // use Timestamps
		ClientBonus        // has one client bonus
		Phone       string `gorm:"type:VARCHAR(11);UNIQUE_INDEX;NOT NULL"`
	}
)

var (
	// Make sure the type satisfies manager interface.
	_ OrmManager = (*Client)(nil)
)

// NewClient constructs a new entity.
func NewClient(orm *gorm.DB) *Client {
	e := &Client{}

	// set entity properties and relations
	e.PrimaryKey = primaryKey(orm, e)

	return e
}

// TableName returns database table name.
func (*Client) TableName() string {
	return "clients"
}

// Create satisfies entity interface.
func (e *Client) Create() error {
	return e.DB.FirstOrCreate(e).Error
}

// Read satisfies entity interface.
func (e *Client) Read() error {
	return e.DB.FirstOrCreate(e).Error
}

// Update satisfies entity interface.
func (e *Client) Update() error {
	return e.DB.Updates(e).Error
}

// Delete satisfies entity interface.
func (e *Client) Delete() error {
	return e.DB.Delete(e).Error
}
