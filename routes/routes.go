package routes

import (
	"BelajarAPI/config"
	activity "BelajarAPI/features/activity"
	user "BelajarAPI/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController, ac activity.ActivityController) {
	userRoute(c, ctl)
	activityRoute(c, ac)
}

func userRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/users", ctl.Add())
	c.POST("/login", ctl.Login())
	c.GET("/profile", ctl.Profile(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}

func activityRoute(c *echo.Echo, ac activity.ActivityController) {
	c.POST("/activities", ac.Add(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/activities", ac.ShowMyActivity(), echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/activities/:id", ac.Update(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/activities/:id", ac.Delete(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}