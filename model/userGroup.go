package model

import (
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);unique_index"`
}

func InitUserGroup(db *gorm.DB) {
	var count int64
	db.Model(&UserGroup{}).Where("name = ?", "admin").Count(&count)
	var admin = UserGroup{Name: "admin"}
	if count == 0 {
		db.Create(&admin)
	}

	db.Model(&UserGroup{}).Where("name = ?", "cashier").Count(&count)
	var cashier = UserGroup{Name: "cashier"}
	if count == 0 {
		db.Create(&cashier)
	}

	db.Model(&UserGroup{}).Where("name = ?", "waiter").Count(&count)
	var waiter = UserGroup{Name: "waiter"}
	if count == 0 {
		db.Create(&waiter)
	}

	db.Model(&UserGroup{}).Where("name = ?", "kitchen").Count(&count)
	var kitchen = UserGroup{Name: "kitchen"}
	if count == 0 {
		db.Create(&kitchen)
	}
}
