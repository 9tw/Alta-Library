package controller

import "main.go/model"

type BooksController struct {
	Model model.BooksModel
}

func (bc BooksController) GetAll() ([]model.Res, error) {
	res, err := bc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BooksController) GetUnBorrow() ([]model.Res, error) {
	res, err := bc.Model.GetUnBorrow()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BooksController) OwnBook(id int) ([]model.Res, error) {
	res, err := bc.Model.OwnBook(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BooksController) Search(key string) ([]model.Res, error) {
	res, err := bc.Model.Search(key)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BooksController) AddBook(data model.Books) error {
	err := bc.Model.AddBook(data)
	if err != nil {
		return nil
	}
	return nil
}

func (bc BooksController) UpdateBook(data model.Books) error {
	err := bc.Model.UpdateBook(data)
	if err != nil {
		return nil
	}
	return nil
}

func (bc BooksController) DeleteBook(book_id int) error {
	err := bc.Model.DeleteBook(book_id)
	if err != nil {
		return nil
	}
	return nil
}
