// Copyright © 2020 The EVEN Lab Team

package entity

import (
	"evenlab/go-priority-api/app/http/request"
)

type (
	// Client describes of client entity.
	Client struct {
		// Use PrimaryKey.
		PrimaryKey

		// Use Timestamps.
		Timestamps

		// Client's properties.
		Phone string `gorm:"Type:VARCHAR(11);UNIQUE_INDEX;NOT NULL"`
	}
)

var (
	// Make sure the Client type satisfies the entity interface.
	_ Iface = (*Client)(nil)
)

// NewClient returns a new Client type pointer.
func NewClient(form *request.ClientForm) *Client {
	// Initialize a client entity.
	client := &Client{}

	// Fill entity values.
	client.PrimaryKey = primaryKey(client)
	client.Phone = form.Client.PhoneNumber.Value()

	return client
}

// Set Client's entity database table name.
func (*Client) TableName() string {
	return "clients"
}

// Create satisfies entity interface.
func (c *Client) Create() error {
	return c.DB.FirstOrCreate(c).Error
}

// Read satisfies entity interface.
func (c *Client) Read() error {
	return c.DB.First(c).Error
}

// Update satisfies entity interface.
func (c *Client) Update() error {
	return c.DB.Updates(c).Error
}

// Delete satisfies entity interface.
func (c *Client) Delete() error {
	return c.DB.Delete(c).Error
}
