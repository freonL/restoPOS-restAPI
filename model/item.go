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

func InitItem(db *gorm.DB) {
	var count int64
	db.Model(&Item{}).Count(&count)
	if count == 0 {
		items := []Item{
			{
				Name:        "Nasi Goreng",
				Description: "Indoensian Fried Rice",
				Price:       8.75,
				CategoryID:  1,
				IsActive:    true,
			},
			{
				Name:        "Steam Rice",
				Description: "Steam Rice",
				Price:       3,
				CategoryID:  1,
				IsActive:    true,
			},
			{
				Name:        "Mie Goreng",
				Description: "Indoensian Fried Noodle",
				Price:       8.75,
				CategoryID:  1,
				IsActive:    true,
			},

			{
				Name:        "Ayam Goreng",
				Description: "Indoensian Fried Chicker serve with Tofu and Chili",
				Price:       6.50,
				CategoryID:  2,
				IsActive:    true,
			},

			{
				Name:        "Beef Rendang",
				Description: "Caramelise Beef with Spice & Coconut Milk",
				Price:       7.50,
				CategoryID:  3,
				IsActive:    true,
			},

			{
				Name:        "Gado-Gado",
				Description: "Indonesian Salad with Peanut Sauce",
				Price:       7.20,
				CategoryID:  4,
				IsActive:    true,
			},

			{
				Name:        "Soft Drink",
				Description: "350ml can",
				Price:       3,
				CategoryID:  5,
				IsActive:    true,
			},

			{
				Name:        "Kopi-O",
				Description: "Hot Black Coffee",
				Price:       3,
				CategoryID:  5,
				IsActive:    true,
			},
		}

		db.Create(&items)
	}
}
