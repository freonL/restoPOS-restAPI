package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50)"`
	Description string
	Price       float32
	CategoryID  uint
	Category    ItemCategory `json:"category"`
	IsActive    bool         `gorm:"default:true"`
}

func (e *Item) Disable() {
	e.IsActive = false
}

func (p *Item) Enable() {
	p.IsActive = true
}
