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
	err := um.DB.Table("users").Select("id", "email", "name", "address", "status").Where("email = ? and password = ?", key, Password).Model(&Users{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil

}
func (um UsersModel) AddAccount(newData Users) ([]Res, error) {

	var res []Res
	err := um.DB.Exec("INSERT INTO users (email, name, phone, address, password, status, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)",
		newData.Email, newData.Name, newData.Phone, newData.Address, newData.Password, newData.Status, time.Now(), time.Now()).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil

}
func (um UsersModel) Update(editData Users) ([]Res, error) {
	var res []Res
	err := um.DB.Exec("UPDATE users SET name = ? where phone = ?",
		editData.Name, editData.Phone).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (um UsersModel) UpdatePhone(editNomor Users) ([]Res, error) {
	var res []Res
	err := um.DB.Exec("UPDATE users SET phone = ? where password = ?",
		editNomor.Phone, editNomor.Password).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (um UsersModel) UpdateAlamat(editData Users) ([]Res, error) {
	var res []Res
	err := um.DB.Exec("UPDATE users SET address = ? where password = ?",
		editData.Address, editData.Password).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}


func (um UsersModel) UpdateStatus(editData Users) ([]Res, error) {
	var res []Res
	err := um.DB.Exec("UPDATE users SET status = ? where password = ?",
		editData.Status, editData.Password).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}
