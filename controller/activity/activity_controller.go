package activity

import (
	"BelajarAPI/helper"
	"BelajarAPI/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ActivityController struct {
	Model model.ActivityModel
}

func (ac *ActivityController) AddActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddActivityRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim tidak sesuai", nil))
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {
			var message []string
			for _, val := range err.(validator.ValidationErrors) {
				message = append(message, fmt.Sprintf("%s %s", val.Field(), val.Tag()))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim kurang sesuai", message))
		}

		var newData model.Activity
		newData.UserHp = input.UserHp
		newData.Title = input.Title
		newData.Description = input.Description

		err = ac.Model.AddActivity(newData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}

		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated, "activity berhasil ditambahkan", nil))
	}
}

func (ac *ActivityController) GetActivityByUserHp() echo.HandlerFunc {
	return func(c echo.Context) error {
		userHP := c.Param("userHP")
		activities, err := ac.Model.GetActivityByUserHp(userHP)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", activities))
	}
}

func (ac *ActivityController) UpdateActivityByUserHp() echo.HandlerFunc {
	return func(c echo.Context) error {
		userHP := c.Param("userHP")
		var input UpdateActivityRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim tidak sesuai", nil))
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {
			var message []string
			for _, val := range err.(validator.ValidationErrors) {
				message = append(message, fmt.Sprintf("%s %s", val.Field(), val.Tag()))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim kurang sesuai", message))
		}

		var newData model.Activity
		newData.UserHp = input.UserHp
		newData.Title = input.Title
		newData.Description = input.Description

		err = ac.Model.UpdateActivityByUserHp(userHP, newData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "activity berhasil diupdate", nil))
	}
}

func (ac *ActivityController) DeleteActivityByUserHp() echo.HandlerFunc {
	return func(c echo.Context) error {
		userHP := c.Param("userHP")

		err := ac.Model.DeleteActivityByUserHp(userHP)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "activity berhasil dihapus", nil))
	}
}
