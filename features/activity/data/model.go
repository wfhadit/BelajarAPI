package data

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Pemilik string `gorm:"type:varchar(13);"`
	Title string
	Description string
}