package model

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name  string
	Price uint
}
