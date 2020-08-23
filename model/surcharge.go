package model

import "gorm.io/gorm"

type Surcharge struct {
	gorm.Model
	Name    string
	Percent int
	Order   int
}
