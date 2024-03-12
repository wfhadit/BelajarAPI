package handler

import (
	"BelajarAPI/features/activity"
	"BelajarAPI/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
    s activity.ActivityService
}

// NewHandler membuat instance baru dari ActivityController.
func NewHandler(service activity.ActivityService) activity.ActivityController {
    return &controller{
        s: service,
    }
}

// Add menangani penambahan aktivitas.
func (ct *controller) Add() echo.HandlerFunc {
    return func(c echo.Context) error {
        var input ActivityRequest
        err := c.Bind(&input)
        if err != nil {
            log.Println("error bind data:", err.Error())
            if strings.Contains(err.Error(), "unsupport") {
                return c.JSON(http.StatusUnsupportedMediaType,
                    helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
            }
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        token, ok := c.Get("user").(*jwt.Token)
        if !ok {
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        newActivity, err := ct.s.AddActivity(token, input.Judul, input.Deskripsi)
        if err != nil {
            log.Println("error insert db:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
        }

        return c.JSON(http.StatusCreated,
            helper.ResponseFormat(http.StatusCreated, "berhasil menambahkan kegiatan", newActivity))
    }
}

func (ct *controller) Update() echo.HandlerFunc {
    return func(c echo.Context) error {
        // Mendapatkan id aktivitas dari parameter URL
        idStr := c.Param("id")
        id, err := strconv.ParseUint(idStr, 10, 64)
        if err != nil {
            log.Println("error parsing ID:", err.Error())
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        // Binding data input ke struct ActivityRequest
        var input ActivityRequest
        if err := c.Bind(&input); err != nil {
            log.Println("error bind data:", err.Error())
            if strings.Contains(err.Error(), "unsupported") {
                return c.JSON(http.StatusUnsupportedMediaType,
                    helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
            }
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        // Mendapatkan token pengguna dari context
        token, ok := c.Get("user").(*jwt.Token)
        if !ok {
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        // Memanggil service untuk melakukan pembaruan aktivitas
        updatedActivity, err := ct.s.UpdateActivity(token, uint(id), activity.Activity{
            Judul:     input.Judul,
            Deskripsi: input.Deskripsi,
        })
        if err != nil {
            log.Println("error update activity:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
        }

        // Mengembalikan respons dengan aktivitas yang telah diperbarui
        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "aktivitas berhasil diperbarui", updatedActivity))
    }
}

// Delete menangani penghapusan aktivitas.
func (ct *controller) Delete() echo.HandlerFunc {
    return func(c echo.Context) error {
        // Mendapatkan id aktivitas dari parameter URL
        idStr := c.Param("id")
        id, err := strconv.ParseUint(idStr, 10, 64)
        if err != nil {
            log.Println("error parsing ID:", err.Error())
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        // Mendapatkan token pengguna dari context
        token, ok := c.Get("user").(*jwt.Token)
        if !ok {
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        // Memanggil service untuk menghapus aktivitas
        err = ct.s.DeleteActivity(token, uint(id))
        if err != nil {
            log.Println("error delete activity:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
        }

        // Mengembalikan respons HTTP sukses setelah penghapusan berhasil
        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "aktivitas berhasil dihapus", nil))
    }
}



// ShowMyActivity menampilkan aktivitas pengguna.
func (ct *controller) ShowMyActivity() echo.HandlerFunc {
    return func(c echo.Context) error {
        // Mendapatkan token pengguna dari context
        token, ok := c.Get("user").(*jwt.Token)
        if !ok {
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        // Memanggil service untuk mendapatkan aktivitas pengguna
        activities, err := ct.s.GetActivityByOwner(token)
        if err != nil {
            log.Println("error get user activities:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
        }

        // Mengembalikan respons HTTP dengan aktivitas pengguna
        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "aktivitas pengguna", activities))
    }
}
