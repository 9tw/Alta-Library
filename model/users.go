package model

import (
	// "fmt"
	"fmt"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string
	Name     string
	Phone    string
	Address  string
	Password string
	Status   string
}

type Res struct {
	Email    string
	Name     string
	Phone    string
	Address  string
	Password string
	Status   string
}

type UsersModel struct {
	DB *gorm.DB
}

func (um UsersModel) GetAll() ([]Res, error) {
	var res []Res
	err := um.DB.Table("users").Select("id", "email", "name", "phone", "address", "password", "status").Model(&Users{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

	func (um UsersModel) LoginUser(key, Password string) ([]Res, error) {
		var res []Res
		err := um.DB.Table("users").Select("id", "email", "name", "address", "status").Where("email = ? and password = ?", key, Password ).Model(&Users{}).Find(&res).Error
		if err != nil {
			fmt.Println("error on query", err.Error())
			return nil, err
		}
		return res, nil

	}