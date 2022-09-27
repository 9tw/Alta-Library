package main

import (
	"fmt"
	"os"
	"os/exec"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/controller"
	"main.go/model"
)

func connect() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/library"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Books{})
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var run bool = true
	var input int

	conn, err := connect()
	//migrate(conn)
	booksM := model.BooksModel{conn}
	booksC := controller.BooksController{booksM}
	if err != nil {
		fmt.Println("Cannot connect to DB", err.Error())
	}

	for run {
		fmt.Println("--- Welcome to Alta Library ---")
		fmt.Println("")
		fmt.Println("1. Search Books")
		fmt.Println("2. List All Books")
		fmt.Println("3. Register")
		fmt.Println("4. Login")
		fmt.Println("9. Exit")
		fmt.Println("")
		fmt.Print("Enter Input: ")
		fmt.Scan(&input)

		switch input {
		case 1:
			var keyword string
			fmt.Println("--- Search Books in Alta Library ---")
			fmt.Println("")
			fmt.Print("Keyword: ")
			fmt.Scan(&keyword)
			res, err := booksC.Search(keyword)
			if err != nil {
				fmt.Println("Some error on get", err.Error())
			}
			fmt.Println(res)
		case 2:
			fmt.Println("--- List All Books in Alta Library ---")
			fmt.Println("")
			res, err := booksC.GetAll()
			if err != nil {
				fmt.Println("Some error on get", err.Error())
			}
			fmt.Println(res)
		case 3:
			var email string
			var name string
			var phone string
			var address string
			var password string
			fmt.Println("--- Register to Alta Library ---")
			fmt.Println("")
			fmt.Print("Email: ")
			fmt.Scan(&email)
			fmt.Print("Name: ")
			fmt.Scan(&name)
			fmt.Print("Phone: ")
			fmt.Scan(&phone)
			fmt.Print("Address: ")
			fmt.Scan(&address)
			fmt.Print("Password: ")
			fmt.Scan(&password)
		case 4:
			var email string
			var password string
			fmt.Println("--- Login to Alta Library ---")
			fmt.Println("")
			fmt.Print("Email: ")
			fmt.Scan(&email)
			fmt.Print("Password: ")
			fmt.Scan(&password)
		case 9:
			clear()
			run = false
			fmt.Println("Bye, see u again.")
		}
	}
}
