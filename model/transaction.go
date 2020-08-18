package model

import (
	"gorm.io/gorm"
)

type TransHeader struct {
	gorm.Model
	EntryUserID uint
	EntryUser   User
	Customer    string
	Status      int
}

type TransDetail struct {
	gorm.Model
	HeaderID     uint
	Header       TransHeader
	ItemID       uint
	Item         Item
	QtyOrdered   uint
	QtyDelivered uint
	Note         string
	Status       int
}
