package data

import (
	"errors"
	"fmt"
	"log"
	"pvg/domain"

	_bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) Insert(newUser domain.User) (row int, err error) {
	user := FromModel(newUser)
	passwordHashed, errorHash := _bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorHash != nil {
		fmt.Println("Error hash", errorHash.Error())
	}
	user.Password = string(passwordHashed)
	resultCreate := ud.db.Create(&user)
	if resultCreate.Error != nil {
		return 0, resultCreate.Error
	}
	if resultCreate.RowsAffected != 1 {
		return 0, errors.New("failed to insert data, your email is already registered")
	}
	return int(resultCreate.RowsAffected), nil
}

func (ud *userData) DeleteData(userID int) (row int, err error) {
	res := ud.db.Delete(&User{}, userID)
	if res.Error != nil {
		log.Println("cannot delete data", res.Error.Error())
		return 0, res.Error
	}
	if res.RowsAffected < 1 {
		log.Println("no data deleted", res.Error.Error())
		return 0, errors.New("failed to delete data ")
	}
	return int(res.RowsAffected), nil
}

func (ud *userData) GetAllUser(limit, offset int) (data []domain.User, err error) {
	dataUser := []User{}
	res := ud.db.Model(&User{}).Limit(limit).Offset(offset).Find(&dataUser)
	if res.Error != nil {
		return []domain.User{}, nil
	}
	return ParseToArr(dataUser), nil
}

func (ud *userData) GetSpecific(userID int) (domain.User, error) {
	var tmp User
	err := ud.db.Where("ID = ?", userID).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return domain.User{}, err
	}

	return tmp.ToModel(), nil
}

func (ud *userData) UpdateData(data map[string]interface{}, idParam int) (row int, err error) {
	res := ud.db.Model(&User{}).Where("id = ?", idParam).Updates(data)
	if res.Error != nil {
		return 0, res.Error
	}
	if res.RowsAffected != 1 {
		return 0, errors.New("failed update data")
	}
	return int(res.RowsAffected), nil
}
