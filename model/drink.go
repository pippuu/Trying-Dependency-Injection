package model

import "gorm.io/gorm"

type Drink struct {
	gorm.Model
	Name    string
	Price   uint
	Topping []*Topping `gorm:"many2many:drink_topping;"`
}
