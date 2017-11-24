package migrations

import (
	"log"

	"github.com/jinzhu/gorm"

	"lab/go-rest-api/app/entity"
)

func AutoMigrate(orm *gorm.DB) {
	log.Println("Start auto-migration...")

	// append "ENGINE=InnoDB" to the SQL statement when creating tables
	orm.Set("gorm:table_options", "ENGINE=InnoDB")

	// try auto-migrate database tables with specified entities
	if err := orm.AutoMigrate(
		&entity.Client{},
		&entity.ClientBonus{},
		&entity.Organization{},
		&entity.BusinessUnit{},
		&entity.WorkPlace{},
		&entity.Document{},
		&entity.Payment{},
		&entity.Position{},
	).Error; err != nil {
		log.Fatalf("automigrate: %+v\n", err)
	}
}
