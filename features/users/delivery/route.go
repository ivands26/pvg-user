package delivery

import (
	"pvg/domain"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, du domain.UserHandler) {
	e.POST("/users", du.InsertUser())
	e.DELETE("/users/:id", du.DeleteById())
	e.GET(("/users"), du.GetAllUser())
	e.GET(("/users/:id"), du.GetUserByID())
	e.PUT(("/users"), du.UpdateUser())

}
