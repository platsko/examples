// Copyright © 2020 The EVEN Lab Team

package entity

import "time"

type (
	// Timestamps is a struct that may extends of entity definition,
	// which could be embedded in other entities.
	//
	// Including fields:
	// `CreatedAt`
	// `UpdatedAt`
	// `DeletedAt`
	Timestamps struct {
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time `gorm:"INDEX"`
	}
)
