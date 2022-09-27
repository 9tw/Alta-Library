package controller

import "main.go/model"

type BooksController struct {
	Model model.BooksModel
}

func (bc BooksController) GetAll() ([]model.Books, error) {
	res, err := bc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BooksController) Search(key string) ([]model.Books, error) {
	res, err := bc.Model.Search(key)
	if err != nil {
		return nil, err
	}
	return res, nil
}
