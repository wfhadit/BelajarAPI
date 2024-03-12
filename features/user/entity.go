package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Add() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
	Profile(token *jwt.Token) (User, error)
}

type UserModel interface {
	AddUser(newData User) error
	Login(hp string) (User, error)
	GetUserByHP(hp string) (User, error)
}

type User struct {
	Nama 		string
	Hp 			string
	Password 	string
}

type Login struct {
	Hp 			string `validate:"required,min=10,max=13,numeric"`
	Password 	string `validate:"required,alphanum,min=8"`
}

type Register struct {
	Nama 		string `validate:"required,alpha"`
	Hp 			string `validate:"required,min=10,max=13,numeric"`
	Password 	string `validate:"required,alphanum,min=8"`
}