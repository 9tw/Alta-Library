package controller

import "main.go/model"

type BorrowsController struct {
	Model model.BorrowsModel
}

func (bc BorrowsController) BorrowBook(data model.Borrows) error {
	err := bc.Model.Borrow(data)
	if err != nil {
		return nil
	}
	return nil
}

func (bc BorrowsController) ListBorrow(id int) ([]model.Result, error) {
	res, err := bc.Model.ListBorrow(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BorrowsController) ReturnBook(book_id int, user_id int) error {
	err := bc.Model.Return(book_id, user_id)
	if err != nil {
		return nil
	}
	return nil
}
