package services

import (
	"BelajarAPI/features/activity"
	"BelajarAPI/helper"
	"BelajarAPI/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m activity.ActivityModel
	v *validator.Validate
}

func NewActivityService(model activity.ActivityModel) activity.ActivityService {
	return &service{
		m: model,
		v: validator.New(),
	}
}

func (s *service) AddActivity(pemilik *jwt.Token, judulBaru string, deskripsiBaru string) (activity.Activity, error) {
	hp := middlewares.DecodeToken(pemilik)
	if hp == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return activity.Activity{}, errors.New("data tidak valid")
	}

	err := s.v.Var(judulBaru, "required")
	if err != nil {
		log.Println("error validasi judul", err.Error())
		return activity.Activity{}, err
	}

	err = s.v.Var(deskripsiBaru, "required")
	if err != nil {
		log.Println("error validasi deskripsi", err.Error())
		return activity.Activity{}, err
	}

	result, err := s.m.AddActivity(hp, judulBaru, deskripsiBaru)
	if err != nil {
		return activity.Activity{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}

func (s *service) UpdateActivity(pemilik *jwt.Token, activityID uint, data activity.Activity) (activity.Activity, error) {
	hp := middlewares.DecodeToken(pemilik)
	if hp == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return activity.Activity{}, errors.New("data tidak valid")
	}

	// Melakukan validasi terhadap data aktivitas yang akan diupdate
	err := s.v.Struct(data)
	if err != nil {
		log.Println("error validasi aktivitas", err.Error())
		return activity.Activity{}, err
	}

	// Memanggil method UpdateActivity dari model untuk melakukan update
	result, err := s.m.UpdateActivity(hp, activityID, data)
	if err != nil {
		return activity.Activity{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}

// DeleteActivity menghapus aktivitas berdasarkan ID dan pemiliknya.
func (s *service) DeleteActivity(pemilik *jwt.Token, activityID uint) error {
    hp := middlewares.DecodeToken(pemilik)
    if hp == "" {
        log.Println("error decode token:", "token tidak ditemukan")
        return errors.New("data tidak valid")
    }

    // Memanggil method DeleteActivity dari model untuk menghapus aktivitas
    err := s.m.DeleteActivity(activityID)
    if err != nil {
        return errors.New(helper.ServerGeneralError)
    }

    return nil
}


func (s *service) GetActivityByOwner(pemilik *jwt.Token) ([]activity.Activity, error) {
	hp := middlewares.DecodeToken(pemilik)
	if hp == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return nil, errors.New("data tidak valid")
	}

	// Memanggil method GetActivityByOwner dari model untuk mendapatkan aktivitas berdasarkan pemilik
	activities, err := s.m.GetActivityByOwner(hp)
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}

	return activities, nil
}
