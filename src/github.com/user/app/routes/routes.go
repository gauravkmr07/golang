package routes

import (
	"github.com/labstack/echo"
	"github.com/user/app/controllers"
)

// LoadRoutes :: It will be used to load all routes used in application
func LoadRoutes(e *echo.Echo) {
	e.GET("/employees", controllers.ListEmployee)
	e.POST("/employees", controllers.CreateEmployee)
	e.PUT("/employees/:id", controllers.UpdateEmployee)
}
