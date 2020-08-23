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

func InitItemCategory(db *gorm.DB) {
	var count int64
	db.Model(&ItemCategory{}).Count(&count)
	if count == 0 {
		categories := []ItemCategory{
			{
				Name:     "Rice & Noodle",
				IsActive: true,
			},
			{
				Name:     "Chicken & Fish",
				IsActive: true,
			},
			{
				Name:     "Meat",
				IsActive: true,
			},
			{
				Name:     "Vegetable",
				IsActive: true,
			},
			{
				Name:     "Beverage",
				IsActive: true,
			},
		}
		db.Create(&categories)
	}

}
