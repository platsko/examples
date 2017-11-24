package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"lab/go-rest-api/config"
)

// mySqlConnect creates connection to MySql server.
func mySqlConnect(db config.DbConfig) (*gorm.DB, error) {
	// connection string format expects six parameters:
	// USER; PASS; HOST; PORT; DBNAME; CHARSET
	// Example: *USER*:*PASS*@tcp([*HOST*]:*PORT*)/*DBNAME*?charset=*CHARSET*
	format := "%s:%s@tcp([%s]:%d)/%s?charset=%s&parseTime=True"
	dsn := fmt.Sprintf(format, db.User, db.Pass, db.Host, db.Port, db.Name, db.Code)

	log.Printf("Connect to MySQL DSN: %s\n", dsn)

	return gorm.Open("mysql", dsn)
}
