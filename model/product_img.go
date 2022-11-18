package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ProductId int `gorm:"not null"`
	ImgPath   string
}
