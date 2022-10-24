package domain

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID          int
	Username    string
	Password    string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Birthday    string
}

//logic
type UserUseCase interface {
	AddUser(newUser User) (row int, err error)
	DeleteCase(userID int) (row int, err error)
	GetAllUserCase(limit, offset int) (data []User, err error)
	GetProfile(id int) (User, error)
	UpdateCase(input User, idParam int) (row int, err error)
}

//query
type UserData interface {
	Insert(newUser User) (row int, err error)
	DeleteData(userID int) (row int, err error)
	GetAllUser(limit, offset int) (data []User, err error)
	GetSpecific(userID int) (User, error)
	UpdateData(data map[string]interface{}, idParam int) (row int, err error)
}

//handler
type UserHandler interface {
	InsertUser() echo.HandlerFunc
	DeleteById() echo.HandlerFunc
	GetAllUser() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
}
