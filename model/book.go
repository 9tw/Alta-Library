package model

import (
	"fmt"

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

type BooksModel struct {
	DB *gorm.DB
}

func (bm BooksModel) GetAll() ([]Books, error) {
	var res []Books
	err := bm.DB.Table("books").Select("id", "title", "isbn", "author", "image", "status", "user_id").Model(&Books{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm BooksModel) Search(key string) ([]Books, error) {
	var res []Books
	err := bm.DB.Table("books").Select("id", "title", "isbn", "author", "image", "status", "user_id").Where("title LIKE ?", "%"+key+"%").Model(&Books{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}
