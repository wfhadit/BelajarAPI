package routes

import (
	"BelajarAPI/config"
	"BelajarAPI/controller/activity"
	"BelajarAPI/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController, ctr activity.ActivityController) {
	c.POST("/users", ctl.Register())
	c.POST("/login", ctl.Login())

	// Activities endpoints with JWT middleware
	c.POST("/activity", ctr.AddActivity(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/activities/:userHP", ctr.GetActivityByUserHp(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/activity/:userHP", ctr.UpdateActivityByUserHp(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/activity/:userHP", ctr.DeleteActivityByUserHp(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
