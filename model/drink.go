package model

import "gorm.io/gorm"

type Drink struct {
	gorm.Model
	Name    string
	Topping string
	Price   uint
}
