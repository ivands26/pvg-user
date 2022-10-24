package data

import (
	"pvg/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Password    string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Birthday    string
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:          int(u.ID),
		Username:    u.Username,
		Password:    u.Password,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
		Birthday:    u.Birthday,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User
	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Username = data.Username
	res.FirstName = data.FirstName
	res.PhoneNumber = data.PhoneNumber
	res.Birthday = data.Birthday
	res.Email = data.Email
	res.LastName = data.LastName
	return res
}
