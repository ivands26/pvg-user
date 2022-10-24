package delivery

import (
	"pvg/domain"
)

type InsertFormat struct {
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	FirstName   string `json:"firstname" form:"firstname"`
	LastName    string `json:"lastname" form:"lastname"`
	PhoneNumber string `json:"phonenumber" form:"phonenumber"`
	Email       string `json:"email" form:"email"`
	Birthday    string `json:"birthday" form:"birthday"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		Username:    i.Username,
		FirstName:   i.FirstName,
		Password:    i.Password,
		PhoneNumber: i.PhoneNumber,
		LastName:    i.LastName,
		Email:       i.Email,
		Birthday:    i.Birthday,
	}
}
