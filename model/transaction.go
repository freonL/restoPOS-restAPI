package model

import (
	"gorm.io/gorm"
)

type TransHeader struct {
	gorm.Model
	EntryUserID  uint
	EntryUser    User
	UpdateUserID uint
	UpdateUser   User
	Customer     string

	Details    []TransDetail    `json:"details" gorm:"foreignkey:TransHeaderID"`
	Surcharges []TransSurcharge `json:"surcharges" gorm:"foreignkey:TransHeaderID"`

	GrossTotal float32
	NetTotal   float32
	Payment    int
	Status     int
}

type Qty struct {
	Orded   uint
	Deliver uint
}

type TransDetail struct {
	gorm.Model
	TransHeaderID uint `json:"-"`
	ItemID        uint
	Item          Item
	Qty           Qty `json:"quantity" gorm:"embedded"`
	Note          string
	UnitPrice     float32
	Subtotal      float32
	Status        int
}

type TransSurcharge struct {
	gorm.Model
	TransHeaderID uint
	SurchargeID   uint
	Surcharge     Surcharge
	BaseValue     float32
	ResultValue   float32
}
