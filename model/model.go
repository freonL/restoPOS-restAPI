package model

import (
	"gorm.io/gorm"
)

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&ItemCategory{})
	db.AutoMigrate(&Item{})

	db.AutoMigrate(&UserGroup{})
	InitUserGroup(db)
	db.AutoMigrate(&User{})
	InitUser(db)
	db.AutoMigrate(&TransHeader{})
	db.AutoMigrate(&TransDetail{})
	return db
}
