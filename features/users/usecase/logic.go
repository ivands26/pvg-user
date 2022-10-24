package usecase

import (
	"errors"
	"fmt"
	"log"
	"pvg/domain"

	_bcrypt "golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type userUseCase struct {
	userData domain.UserData
}

func New(ud domain.UserData) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
	}
}

func (uc *userUseCase) AddUser(newUser domain.User) (row int, err error) {
	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" || newUser.PhoneNumber == "" || newUser.Birthday == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}

	row, err = uc.userData.Insert(newUser)
	return row, err
}

func (uc *userUseCase) DeleteCase(userID int) (row int, err error) {
	row, err = uc.userData.DeleteData(userID)
	return row, err
}

func (uc *userUseCase) GetAllUserCase(limit, offset int) (data []domain.User, err error) {
	res, err := uc.userData.GetAllUser(limit, offset)
	return res, err
}

func (uc *userUseCase) GetProfile(id int) (domain.User, error) {
	data, err := uc.userData.GetSpecific(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}

func (uc *userUseCase) UpdateCase(input domain.User, idParam int) (row int, err error) {
	userReq := map[string]interface{}{}
	if input.Username != "" {
		userReq["username"] = input.Username
	}
	if input.Email != "" {
		userReq["email"] = input.Email
	}
	if input.PhoneNumber != "" {
		userReq["phonenumber"] = input.PhoneNumber
	}
	if input.FirstName != "" {
		userReq["firstname"] = input.FirstName
	}
	if input.LastName != "" {
		userReq["lastname"] = input.LastName
	}
	if input.Birthday != "" {
		userReq["birthday"] = input.Birthday
	}
	if input.Password != "" {
		passwordHashed, errorHash := _bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		if errorHash != nil {
			fmt.Println("Error hash", errorHash.Error())
		}
		userReq["password"] = string(passwordHashed)
	}
	if input.Username == "" && input.Email == "" && input.PhoneNumber == "" && input.LastName == "" && input.Birthday == "" && input.FirstName == "" {
		return 404, err
	}
	row, err = uc.userData.UpdateData(userReq, idParam)
	return row, err
}
