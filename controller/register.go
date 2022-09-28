package controller

import "main.go/model"

type RegisterController struct {
	Model model.UsersModel
}

func (rc RegisterController) Register(data model.Users) ([]model.Res, error) {
	res, err := rc.Model.AddAccount(data)
	if err != nil {
		return nil, err
	}
	return res, nil

}