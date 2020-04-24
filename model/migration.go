package model

import "github.com/jinzhu/gorm"

func DBMigrationAccount(db *gorm.DB, dbs interface{}) *gorm.DB {
	db.AutoMigrate(dbs)
	return db
}
