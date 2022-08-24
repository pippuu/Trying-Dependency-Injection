package model

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	Name    string
	Age     uint
	Address string
}
