package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Title  string
	ISBN   string
	Author string
	Image  string
	Status int
	UserID int
}

type Res struct {
	ID     int
	Title  string
	ISBN   string
	Author string
	Image  string
	Info   string
	Name   string
}

type BooksModel struct {
	DB *gorm.DB
}

func (bm BooksModel) GetAll() ([]Res, error) {
	var res []Res
	err := bm.DB.Table("books").Select("books.id", "books.title", "books.isbn", "books.author", "books.image", "statb.info", "users.name").Joins("join users on users.id=books.user_id").Joins("join statb on statb.id=books.status").Where("books.deleted_at IS NULL").Model(&Res{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm BooksModel) GetUnBorrow() ([]Res, error) {
	var res []Res
	err := bm.DB.Table("books").Select("books.id", "books.title", "books.isbn", "books.author", "books.image", "statb.info", "users.name").Joins("join users on users.id=books.user_id").Joins("join statb on statb.id=books.status").Where("books.deleted_at IS NULL AND books.status = 0").Model(&Res{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm BooksModel) Search(key string) ([]Res, error) {
	var res []Res
	err := bm.DB.Table("books").Select("books.id", "books.title", "books.isbn", "books.author", "books.image", "statb.info", "users.name").Joins("join users on users.id=books.user_id").Joins("join statb on statb.id=books.status").Where("books.title LIKE ? AND books.deleted_at IS NULL", "%"+key+"%").Model(&Books{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm BooksModel) AddBook(newData Books) error {
	err := bm.DB.Exec("INSERT INTO books (id, created_at, updated_at, deleted_at, title, isbn, author, image, status, user_id) values (?,?,?,?,?,?,?,?,?,?)",
		nil, time.Now(), nil, nil, newData.Title, newData.ISBN, newData.Author, newData.Image, 0, newData.UserID).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil
	}
	return nil
}

func (bm BooksModel) UpdateBook(editData Books) error {
	err := bm.DB.Exec("UPDATE books SET updated_at = ?, title = ?, isbn = ?, author = ?, image = ?, status = ?, user_id = ? WHERE id = ?",
		time.Now(), editData.Title, editData.ISBN, editData.Author, editData.Image, editData.Status, editData.UserID, editData.ID).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return nil
}

func (bm BooksModel) DeleteBook(book_id int) error {
	err := bm.DB.Exec("UPDATE books SET deleted_at = ? WHERE id = ?", time.Now(), book_id).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil
	}
	return nil
}
