package data

import (
	"BelajarAPI/features/activity"
	"errors"

	"gorm.io/gorm"
)

type model struct {
    connection *gorm.DB
}

// New membuat instance baru dari ActivityModel.
func New(db *gorm.DB) activity.ActivityModel {
    return &model{
        connection: db,
    }
}

// AddActivity menambahkan aktivitas baru ke dalam database.
func (am *model) AddActivity(pemilik string, judulBaru string, deskripsiBaru string) (activity.Activity, error) {
    var inputProcess = Activity{Judul: judulBaru, Deskripsi: deskripsiBaru, Pemilik: pemilik}
    if err := am.connection.Create(&inputProcess).Error; err != nil {
        return activity.Activity{}, err
    }

    return activity.Activity{Judul: inputProcess.Judul, Deskripsi: inputProcess.Deskripsi}, nil
}

// UpdateActivity mengupdate aktivitas berdasarkan pemilik, ID aktivitas, dan data aktivitas yang baru.
func (am *model) UpdateActivity(pemilik string, activityID uint, data activity.Activity) (activity.Activity, error) {
    var qry = am.connection.Where("pemilik = ? AND id = ?", pemilik, activityID).Updates(data)
    if err := qry.Error; err != nil {
        return activity.Activity{}, err
    }

    if qry.RowsAffected < 1 {
        return activity.Activity{}, errors.New("no data affected")
    }

    return data, nil
}

// DeleteActivity menghapus aktivitas berdasarkan ID aktivitas.
func (am *model) DeleteActivity(activityID uint) error {
    result := am.connection.Delete(&Activity{}, activityID)
    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("no data affected")
    }

    return nil
}

// GetActivityByOwner mendapatkan aktivitas berdasarkan pemilik.
func (am *model) GetActivityByOwner(pemilik string) ([]activity.Activity, error) {
    var result []activity.Activity
    if err := am.connection.Where("pemilik = ?", pemilik).Find(&result).Error; err != nil {
        return nil, err
    }

    return result, nil
}
