package model

import (
	// "fmt"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string
	Name     string
	Phone    string
	Address  string
	Password string
	Status   int
}

type Ress struct {
	Email    string
	Name     string
	Phone    string
	Address  string
	Password string
	Status   int
}

type UsersModel struct {
	DB *gorm.DB
}

func (um UsersModel) GetAll() ([]Ress, error) {
	var res []Ress
	err := um.DB.Table("users").Select("id", "email", "name", "phone", "address", "password", "status").Model(&Users{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (um UsersModel) LoginUser(key, Password string) ([]Ress, error) {
	var res []Ress
	err := um.DB.Table("users").Select("id", "email", "name", "address", "status").Where("email = ? and password = ?", key, Password).Model(&Users{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil

}

func (um UsersModel) AddAccount(newData Users) ([]Ress, error) {

	var res []Ress
	err := um.DB.Exec("INSERT INTO users (email, name, phone, address, password, status, created_at) values (?, ?, ?, ?, ?, ?, ?)",
		newData.Email, newData.Name, newData.Phone, newData.Address, newData.Password, 1, time.Now()).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}
