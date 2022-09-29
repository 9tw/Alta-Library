package controller

import "main.go/model"

type EditProfileController struct {
	Model model.UsersModel
}

func (ep EditProfileController) UpdateProfile(updateData model.Users) ([]model.Res, error) {
	res, err := ep.Model.Update(updateData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ep EditProfileController) UbahPhone(updateNomor model.Users) ([]model.Res, error) {
	res, err := ep.Model.UpdatePhone(updateNomor)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ep EditProfileController) UbahAlamat(updateData model.Users) ([]model.Res, error) {
	res, err := ep.Model.UpdateAlamat(updateData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ep EditProfileController) UbahStatus(updateData model.Users) ([]model.Res, error) {
	res, err := ep.Model.UpdateStatus(updateData)
	if err != nil {
		return nil, err
	}
	return res, nil
}
