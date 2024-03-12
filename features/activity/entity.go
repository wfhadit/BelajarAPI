package activity

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// ActivityController adalah interface yang mendefinisikan method yang diperlukan untuk mengelola aktivitas dalam controller.
type ActivityController interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	ShowMyActivity() echo.HandlerFunc
}

// ActivityModel adalah interface yang mendefinisikan method yang diperlukan untuk mengelola aktivitas dalam model.
type ActivityModel interface {
	AddActivity(pemilik string, judulBaru string, deskripsiBaru string) (Activity, error)
	UpdateActivity(pemilik string, activityID uint, data Activity) (Activity, error)
	DeleteActivity(activityID uint) error
	GetActivityByOwner(pemilik string) ([]Activity, error)
}

type ActivityService interface {
    AddActivity(pemilik *jwt.Token, judulBaru string, deskripsiBaru string) (Activity, error)
    UpdateActivity(pemilik *jwt.Token, activityID uint, data Activity) (Activity, error)
    DeleteActivity(pemilik *jwt.Token, activityID uint) error
    GetActivityByOwner(pemilik *jwt.Token) ([]Activity, error)
}


// Activity adalah struktur yang mewakili data aktivitas.
type Activity struct {
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
}
