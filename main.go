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
	dsn := "root:@tcp(127.0.0.1:3306)/library?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Books{})
	db.AutoMigrate(&model.Borrows{})
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
	booksM := model.BooksModel{conn}
	booksC := controller.BooksController{booksM}
	borrowsM := model.BorrowsModel{conn}
	borrowsC := controller.BorrowsController{borrowsM}
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
		fmt.Println("0. Exit")
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
			fmt.Println("{ID | Title | ISBN | Author | Image | Status | Owner}")
			for i := 0; i < len(res); i++ {
				fmt.Println(res[i])
			}
		case 2:
			fmt.Println("--- List All Books in Alta Library ---")
			fmt.Println("")
			res, err := booksC.GetAll()
			if err != nil {
				fmt.Println("Some error on get", err.Error())
			}
			fmt.Println("{ID | Title | ISBN | Author | Image | Status | Owner}")
			for i := 0; i < len(res); i++ {
				fmt.Println(res[i])
			}
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
			res, err := usersC.Register(model.Users{
				Model:    gorm.Model{},
				Email:    email,
				Name:     name,
				Phone:    phone,
				Address:  address,
				Password: password,
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
			if len(res) > 0 {
				var menu bool = true
				for menu {
					fmt.Println("===========================================")
					fmt.Println("\tWelcome to Alta Library")
					fmt.Println("===========================================")
					fmt.Println("")
					fmt.Println("1. Edit profile")
					fmt.Println("2. Non Aktif akun")
					fmt.Println("3. Pinjam Buku")
					fmt.Println("4. Liat Buku Yang Dipinjam")
					fmt.Println("5. Tambah Buku")
					fmt.Println("6. Edit Buku")
					fmt.Println("7. Hapus Buku")
					fmt.Println("8. Lihat Semua Buku")
					fmt.Println("9. Cari Buku")
					fmt.Println("0. Log Out")
					fmt.Println("")
					fmt.Print("Enter Input: ")
					fmt.Scan(&input)

					switch input {
					case 1:
						fmt.Println("==============================")
						fmt.Println("\tEDIT PROFILE")
						fmt.Println("==============================")

					case 2:
						fmt.Println("================================")
						fmt.Println("\tNON AKTIF AKUN")
						fmt.Println("================================")

					case 3:
						fmt.Println("============================")
						fmt.Println("\tPINJAM BUKU")
						fmt.Println("============================")
						fmt.Println("")
						fmt.Println("Buku yang belum dipinjam:")
						res, err := booksC.GetUnBorrow()
						if err != nil {
							fmt.Println("Some error on get", err.Error())
						}
						fmt.Println("{ID | Title | ISBN | Author | Image | Status | Owner}")
						for i := 0; i < len(res); i++ {
							fmt.Println(res[i])
						}
						var newBorrow model.Borrows
						fmt.Println("")
						fmt.Print("Book ID: ")
						fmt.Scan(&newBorrow.BookID)
						fmt.Print("User ID: ")
						fmt.Scan(&newBorrow.UserID)
						er := borrowsC.BorrowBook(newBorrow)
						if err != nil {
							fmt.Println("Some error on get", er.Error())
						}
						fmt.Println("Borrow Book Success")
					case 4:
						fmt.Println("=========================================")
						fmt.Println("\tLIHAT BUKU YANG DIPINJAM")
						fmt.Println("=========================================")
						var user_id int
						var book_id int
						var ans string
						fmt.Println("")
						fmt.Print("User ID: ")
						fmt.Scan(&user_id)
						res, err := borrowsC.ListBorrow(user_id)
						if err != nil {
							fmt.Println("Some error on get", err.Error())
						}
						fmt.Println("{ID | Title | ISBN | Author | Image | Status | Owner}")
						for i := 0; i < len(res); i++ {
							fmt.Println(res[i])
						}
						if len(res) > 0 {
							fmt.Print("Return Book? [Y/N] ")
							fmt.Scan(&ans)
							if ans == "Y" || ans == "y" {
								fmt.Print("Book ID: ")
								fmt.Scan(&book_id)
								err := borrowsC.ReturnBook(book_id, user_id)
								if err != nil {
									fmt.Println("Some error on get", err.Error())
								}
								fmt.Println("Return Book Success")
							} else {
								break
							}
						} else {
							fmt.Println("You're not borrowing any books.")
						}
					case 5:
						fmt.Println("============================")
						fmt.Println("\tTAMBAH BUKU")
						fmt.Println("============================")
						var newBook model.Books
						fmt.Println("")
						fmt.Print("Title: ")
						fmt.Scan(&newBook.Title)
						fmt.Print("ISBN: ")
						fmt.Scan(&newBook.ISBN)
						fmt.Print("Author: ")
						fmt.Scan(&newBook.Author)
						fmt.Print("Image: ")
						fmt.Scan(&newBook.Image)
						fmt.Print("Owner: ")
						fmt.Scan(&newBook.UserID)
						err := booksC.AddBook(newBook)
						if err != nil {
							fmt.Println("Some error on get", err.Error())
						}
						fmt.Println("Add Book Success")
					case 6:
						fmt.Println("============================")
						fmt.Println("\tEDIT BUKU")
						fmt.Println("============================")
						var editBook model.Books
						fmt.Println("")
						fmt.Print("Book ID: ")
						fmt.Scan(&editBook.ID)
						fmt.Print("Title: ")
						fmt.Scan(&editBook.Title)
						fmt.Print("ISBN: ")
						fmt.Scan(&editBook.ISBN)
						fmt.Print("Author: ")
						fmt.Scan(&editBook.Author)
						fmt.Print("Image: ")
						fmt.Scan(&editBook.Image)
						fmt.Print("Owner: ")
						fmt.Scan(&editBook.UserID)
						err := booksC.UpdateBook(editBook)
						if err != nil {
							fmt.Println("Some error on get", err.Error())
						}
						fmt.Println("Update Book Success")
					case 7:
						fmt.Println("============================")
						fmt.Println("\tHAPUS BUKU")
						fmt.Println("============================")
						var book_id int
						fmt.Println("")
						fmt.Print("Book ID: ")
						fmt.Scan(&book_id)
						err := booksC.DeleteBook(book_id)
						if err != nil {
							fmt.Println("Some error on get", err.Error())
						}
						fmt.Println("Delete Book Success")
					case 8:
						fmt.Println("==============================")
						fmt.Println("\tLIHAT SEMUA BUKU")
						fmt.Println("==============================")
						fmt.Println("")
						res, err := booksC.GetAll()
						if err != nil {
							fmt.Println("Some error on get", err.Error())
						}
						fmt.Println("{ID | Title | ISBN | Author | Image | Status | Owner}")
						for i := 0; i < len(res); i++ {
							fmt.Println(res[i])
						}
					case 9:
						fmt.Println("================================")
						fmt.Println("\tCARI BUKU")
						fmt.Println("================================")
						var keyword string
						fmt.Println("")
						fmt.Print("Keyword: ")
						fmt.Scan(&keyword)
						res, err := booksC.Search(keyword)
						if err != nil {
							fmt.Println("Some error on get", err.Error())
						}
						fmt.Println("{ID | Title | ISBN | Author | Image | Status | Owner}")
						for i := 0; i < len(res); i++ {
							fmt.Println(res[i])
						}
					case 0:
						clear()
						menu = false
						fmt.Println("You're Log Out")
					}
				}
			} else {
				break
			}
		case 0:
			clear()
			run = false
			fmt.Println("Bye, see u again.")
		}
	}
}
