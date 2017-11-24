package entity

import (
	"log"

	"github.com/jinzhu/gorm"
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

// primaryKey returns constructed PrimaryKey type.
func primaryKey(orm *gorm.DB, m interface{}, rel ...interface{}) PrimaryKey {
	db := orm.Model(m)
	if rel != nil {
		for _, r := range rel {
			if err := db.Related(r).Error; err != nil {
				log.Fatalf("entity: %+v relate: %+v", m, err)
			}
		}
	}

	return PrimaryKey{DB: db}
}

// Transaction wraps transaction method of database library
// to avoid import the gorm package into caller package.
func (e *PrimaryKey) Transaction(fn func() error) error {
	return e.DB.Transaction(func(*gorm.DB) error {
		return fn()
	})
}
