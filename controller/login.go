package controller

import (
	"main.go/model"
)

type UsersController struct {
	Model model.UsersModel
}

func (uc UsersController) GetAll() ([]model.Go, error) {
	res, err := uc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UsersController) Search(key, Password string) ([]model.Go, error) {
	res, err := uc.Model.LoginUser(key, Password)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UsersController) Register(newData model.Users) ([]model.Go, error) {
	res, err := uc.Model.AddAccount(newData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UsersController) UpdateProfile(updateData model.Users) ([]model.Go, error) {
	res, err := uc.Model.Update(updateData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UsersController) UbahPhone(updateNomor model.Users) ([]model.Go, error) {
	res, err := uc.Model.UpdatePhone(updateNomor)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UsersController) UbahAlamat(updateData model.Users) ([]model.Go, error) {
	res, err := uc.Model.UpdateAlamat(updateData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UsersController) UbahStatus(updateData model.Users) ([]model.Go, error) {
	res, err := uc.Model.UpdateStatus(updateData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// func (uc UsersController) NonAktifAkun(id model.Users) ([]model.Go, error) {
// 	res, err := uc.Model.NonAktif(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
