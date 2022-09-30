package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Borrows struct {
	gorm.Model
	BookID  int
	UserID  int
	DueDate time.Time
}

type Result struct {
	BookID    int
	Title     string
	Name      string
	CreatedAt time.Time
	DueDate   time.Time
}

type BorrowsModel struct {
	DB *gorm.DB
}

func (bm BorrowsModel) Borrow(newBorrow Borrows) error {
	err := bm.DB.Exec("INSERT INTO borrows (id, created_at, updated_at, deleted_at, book_id, user_id, due_date) values (?,?,?,?,?,?,?)",
		nil, time.Now(), nil, nil, newBorrow.BookID, newBorrow.UserID, (time.Now()).Add(168*time.Hour)).Error
	er := bm.DB.Exec("UPDATE books SET updated_at = ?, status = ? WHERE id = ?", time.Now(), 1, newBorrow.BookID).Error
	if err != nil && er != nil {
		fmt.Println("error on query", err.Error())
		return nil
	}
	return nil
}

func (bm BorrowsModel) ListBorrow(id int) ([]Result, error) {
	var res []Result
	err := bm.DB.Table("borrows").Select("borrows.created_at", "borrows.book_id", "books.title", "users.name", "borrows.due_date").Joins("join books on books.id=borrows.book_id").Joins("join users on users.id=borrows.user_id").Where("borrows.user_id = ? AND borrows.updated_at IS NULL", id).Model(&Result{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm BorrowsModel) Return(book_id, user_id int) error {
	er := bm.DB.Exec("UPDATE books SET updated_at = ?, status = ? WHERE id = ?", time.Now(), 0, book_id).Error
	err := bm.DB.Exec("UPDATE borrows SET updated_at = ? WHERE user_id = ? AND book_id = ? AND updated_at IS NULL", time.Now(), user_id, book_id).Error
	if err != nil && er != nil {
		fmt.Println("error on query", err.Error())
		return nil
	}
	return nil
}
