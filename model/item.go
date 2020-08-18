package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string
	Description string
	Price       float32
	CategoryID  *uint
	Category    *ItemCategory
	IsActive    bool `gorm:"default:true"`
}

func (e *Item) Disable() {
	e.IsActive = false
}

func (p *Item) Enable() {
	p.IsActive = true
}
