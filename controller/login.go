package controller

import "main.go/model"

type UsersController struct {
	Model model.UsersModel
}

func (uc UsersController) GetAll() ([]model.Res, error) {
	res, err := uc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UsersController) Search(key, Password string) ([]model.Res, error) {
	res, err := uc.Model.LoginUser(key, Password)
	if err != nil {
		return nil, err
	}
	return res, nil
}