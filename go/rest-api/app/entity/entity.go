// Copyright © 2020 The EVEN Lab Team

package entity

import (
	"github.com/jinzhu/gorm"

	"evenlab/go-priority-api/database"
)

type (
	// Iface describes entity interface.
	Iface interface {
		// Set entity's database table name.
		TableName() string

		// Create creates a new entity.
		Create() error

		// Read reads an existing entity.
		Read() error

		// Create updates an existing entity.
		Update() error

		// Delete deletes an existing entity.
		Delete() error
	}

	// PrimaryKey is a base definition,
	// which could be embedded in other entities.
	//
	// Including fields:
	// `ID` PRIMARY_KEY
	PrimaryKey struct {
		*gorm.DB `json:"-" gorm:"-"`
		ID       uint64 `gorm:"PRIMARY_KEY"`
	}
)

// primaryKey returns constructed PrimaryKey type.
func primaryKey(entity interface{}) PrimaryKey {
	return PrimaryKey{DB: database.Orm().AutoMigrate(entity)}
}
