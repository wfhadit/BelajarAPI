package main

import (
	"BelajarAPI/config"
	td "BelajarAPI/features/activity/data"
	th "BelajarAPI/features/activity/handler"
	ts "BelajarAPI/features/activity/services"
	"BelajarAPI/features/user/data"
	"BelajarAPI/features/user/handler"
	"BelajarAPI/features/user/services"
	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	activityData := td.New(db)
	activityService := ts.NewActivityService(activityData)
	activityHandler := th.NewHandler(activityService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Initialize routes
	routes.InitRoute(e, userHandler, activityHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
