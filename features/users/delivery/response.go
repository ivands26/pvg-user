package delivery

import (
	"pvg/domain"
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

func ParsePUToArr2(arr []domain.User) []map[string]interface{} {
	var arrmap []map[string]interface{}
	for i := 0; i < len(arr); i++ {
		var res = map[string]interface{}{}
		res["id"] = arr[i].ID
		res["username"] = arr[i].Username
		res["firstname"] = arr[i].FirstName
		res["lastname"] = arr[i].LastName
		res["email"] = arr[i].Email
		res["birthday"] = arr[i].Birthday
		res["phonenumber"] = arr[i].PhoneNumber

		arrmap = append(arrmap, res)
	}
	return arrmap
}

func FromModel(data domain.User) User {
	return User{
		ID:          data.ID,
		Username:    data.Username,
		FirstName:   data.FirstName,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		LastName:    data.LastName,
		Email:       data.Email,
		Birthday:    data.Birthday,
	}
}
