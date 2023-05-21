package routes

import (
	"ogs-create-api/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.GET("/activeUser", controllers.GetActiveUsers)
	e.GET("/userPost", controllers.GetUserPost)
}
