package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"lab/go-rest-api/config"
	"lab/go-rest-api/database/migrations"
)

var (
	// orm keeps gorm database instance.
	orm *gorm.DB
)

// Orm returns database connection.
// If an error occurs, prints error message and terminate.
func Orm() *gorm.DB {
	if err := Open(); err != nil {
		log.Fatalf("Fatal error: %+v\n", err)
	}

	return orm
}

// Close closes database connection.
func Close() {
	if orm != nil {
		_ = orm.Close()
	}
}

// Open opens database connection.
func Open() (err error) {
	if orm == nil {
		// fetch database configuration and
		// choose the database management system
		switch conf := config.Db(); conf.Conn {
		case "mysql":
			orm, err = mySqlConnect(conf)
		default:
			err = fmt.Errorf("unsupported database management system")
		}
		// start auto-migrations
		if err == nil {
			migrations.AutoMigrate(orm)
		}
	}

	return err
}
