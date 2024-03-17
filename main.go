package main

import (
	"BelajarAPI/config"
	"BelajarAPI/controller/activity"
	"BelajarAPI/controller/user"
	"BelajarAPI/model"
	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	userModel := model.UserModel{Connection: db}
	userController := user.UserController{Model: userModel}

	activityModel := model.ActivityModel{Connection: db}
	activityController := activity.ActivityController{Model: activityModel}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Initialize routes
	routes.InitRoute(e, userController, activityController)

	e.Logger.Fatal(e.Start(":1323"))
}
