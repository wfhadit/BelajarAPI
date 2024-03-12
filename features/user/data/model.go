package data

import "BelajarAPI/features/activity/data"

type User struct {
	Nama       string
	Hp         string `gorm:"type:varchar(13);primaryKey"`
	Password   string
	Activities []data.Activity `gorm:"foreignKey:Pemilik;references:Hp"`
}