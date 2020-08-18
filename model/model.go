package model

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string
	Description string
	Price       float32
	CategoryID  *uint
	Category    *ItemCategory
	IsActive    bool
}

type ItemCategory struct {
	gorm.Model
	Name     string
	IsActive bool
}

func (e *Item) Disable() {
	e.IsActive = false
}

func (p *Item) Enable() {
	p.IsActive = true
}

func (e *ItemCategory) Disable() {
	e.IsActive = false
}

func (p *ItemCategory) Enable() {
	p.IsActive = true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&ItemCategory{})
	db.AutoMigrate(&Item{})
	return db
}
