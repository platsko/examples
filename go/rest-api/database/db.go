// Copyright © 2020 The EVEN Lab Team

package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"evenlab/go-priority-api/config"
)

var (
	// orm keeps gorm database instance.
	orm *gorm.DB
)

// Orm returns the database connection.
func Orm() *gorm.DB {
	if err := Open(); err != nil {
		log.Fatalf("Fatal error: %+v\n", err)
	}

	return orm
}

// Close closes the database connection.
func Close() {
	if orm != nil {
		_ = orm.Close()
	}
}

// Open opens the database connection using app config.
func Open() (err error) {
	if orm == nil {
		// Fetch database configuration and
		// choose the database management system.
		switch conf := config.DB(); conf.Conn {
		case "mysql":
			orm, err = mySqlConnect(conf)
		default:
			err = fmt.Errorf("unsupported database management system")
		}
	}

	return err
}
