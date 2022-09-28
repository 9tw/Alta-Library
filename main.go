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
	dsn := "root:@ardhi21091996@tcp(127.0.0.1:3306)/alta_library"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Users{})
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
	migrate(conn)
	usersM := model.UsersModel{conn}
	usersC := controller.UsersController{usersM}
	if err != nil {
		fmt.Println("cannot connect to DB", err.Error())
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
		case 2:
			fmt.Println("--- List All Books in Alta Library ---")
			fmt.Println("")
		case 3:
			var email string
			var name string
			var phone string
			var address string
			var password string
			var status string
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
			fmt.Print("Status: ")
			fmt.Scan(&status)
			res, err := usersC.Register(model.Users{

				Model:    gorm.Model{},
				Email:    email,
				Name:     name,
				Phone:    phone,
				Address:  address,
				Password: password,
				Status:   status,
			})
			//  res, err := usersC.Register(email, name, phone, address, password, status )
			if err != nil {
				fmt.Println("some error on get", err.Error)
			}
			fmt.Println(res)
			fmt.Println("==========================================")
			fmt.Println("ANDA SUKSES MELAKUKAN PENDAFTARAN")
			fmt.Println("SELAMAT MENIKMATI FITUR DAN LAYANAN KAMI")
			fmt.Println("==========================================")
		case 4:
			var email string
			var password string
			fmt.Println("--- Login to Alta Library ---")
			fmt.Println("")
			fmt.Print("Email: ")
			fmt.Scan(&email)
			fmt.Print("Password: ")
			fmt.Scan(&password)
			res, err := usersC.Search(email, password)
			if err != nil {
				fmt.Println("some error on get", err.Error())
			}
			fmt.Println(res)
		case 9:
			clear()
			run = false
			fmt.Println("Bye, see u again.")
		}
	}
}
