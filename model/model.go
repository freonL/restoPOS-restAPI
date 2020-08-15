package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Item struct {
	gorm.Model
	Name     string        `json:"name"`
	Desc     string        `json:"description"`
	Category *ItemCategory `json:"category"`
	Price    float32       `json:"price"`
	IsActive bool          `json:"isActive"`
}

type ItemCategory struct {
	gorm.Model
	Name     string `json:"name"`
	IsActive bool   `json:"isActive,omitempty" bson:",omitempty"`
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
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&ItemCategory{})
	return db
}
