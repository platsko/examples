package entity

import (
	"github.com/jinzhu/gorm"

	"lab/go-rest-api/database"
)

type (
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

// primaryKey returns constructed primary key.
func primaryKey(entity interface{}) PrimaryKey {
	return PrimaryKey{DB: database.Orm().AutoMigrate(entity)}
}
