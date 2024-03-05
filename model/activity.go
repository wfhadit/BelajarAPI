package model

import "gorm.io/gorm"

type Activity struct {
	UserHp      string `json:"user_hp" form:"user_hp" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required,min=10,max=100"`
}

type ActivityModel struct {
	Connection *gorm.DB
}

func (am *ActivityModel) AddActivity(newData Activity) error {
	err := am.Connection.Create(&newData).Error
	if err != nil {
		return err
	}
	return nil
}

func (am *ActivityModel) GetActivityByUserHp(userHP string) ([]Activity, error) {
	var activities []Activity
	if err := am.Connection.Where("user_hp = ?", userHP).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (am *ActivityModel) UpdateActivityByUserHp(userHP string, newData Activity) error {
	if err := am.Connection.Model(&Activity{}).Where("user_hp = ?", userHP).Updates(&newData).Error; err != nil {
		return err
	}
	return nil
}

func (am *ActivityModel) DeleteActivityByUserHp(userHP string) error {
	if err := am.Connection.Delete(&Activity{}, "user_hp = ?", userHP).Error; err != nil {
		return err
	}
	return nil
}
