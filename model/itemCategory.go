package model

import "gorm.io/gorm"

type ItemCategory struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50)"`
	IsActive bool   `gorm:"default:true"`
}

func (e *ItemCategory) Disable() {
	e.IsActive = false
}

func (p *ItemCategory) Enable() {
	p.IsActive = true
}
