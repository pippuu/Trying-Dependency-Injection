package model

import "gorm.io/gorm"

type Topping struct {
	gorm.Model
	Name  string
	Drink []*Drink `gorm:"many2many:drink_topping;"`
}
