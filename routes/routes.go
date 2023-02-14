package routes

import (
	"go-mongodb/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.POST("/create", controllers.Create)
	e.GET("/get", controllers.Get)
	e.DELETE("/delete", controllers.Delete)
	e.PUT("/update", controllers.Update)
}
