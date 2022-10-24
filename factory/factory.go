package factory

import (
	ud "pvg/features/users/data"
	userDelivery "pvg/features/users/delivery"
	us "pvg/features/users/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	useCase := us.New(userData)
	userHandler := userDelivery.New(useCase)
	userDelivery.RouteUser(e, userHandler)
}
