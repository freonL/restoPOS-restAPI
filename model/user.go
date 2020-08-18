package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(32);default:'5f4dcc3b5aa765d61d8327deb882cf99'"`
	GroupID  uint
	Group    UserGroup
	IsActive bool `gorm:"default:true"`
}

func (e *User) Disable() {
	e.IsActive = false
}

func (p *User) Enable() {
	p.IsActive = true
}

func InitUser(db *gorm.DB) {
	var count int64
	db.Model(&User{}).Where("name = ?", "SuperUser").Count(&count)
	var admin = User{
		Name:    "SuperUser",
		GroupID: 1,
	}
	if count == 0 {
		db.Create(&admin)
	}
}
