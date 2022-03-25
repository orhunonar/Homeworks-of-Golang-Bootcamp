package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "admin"
	dbname   = "booksdb"
)

var db *sql.DB
var err error

type BookInfo struct {
	ID         int
	Bookname   string
	Pagenumber int
	Stock      int
	Price      int
	Author     string
}

func InsertBook(data BookInfo) {
	_, err = db.Exec("INSERT INTO books(bookname,pagenumber,stock,price,author) VALUES($1,$2,$3,$4,$5)", data.Bookname, data.Pagenumber, data.Stock, data.Price, data.Author)
	if err != nil {
		log.Fatal(err)

	}
}
func init() {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)

	}

}
func UpdateBook(data BookInfo) {

	_, err = db.Exec("UPDATE books SET bookname=$2,pagenumber=$3,stock=$4,price=$5,author=$6 WHERE id=$1", data.ID, data.Bookname, data.Pagenumber, data.Stock, data.Price, data.Author)
	if err != nil {
		log.Fatal(err)

	}

}

func GetBooks() {

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found!!!")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var books []*BookInfo
	for rows.Next() {
		prd := &BookInfo{}
		err := rows.Scan(&prd.ID, &prd.Bookname, &prd.Pagenumber, &prd.Stock, &prd.Price, &prd.Author)
		if err != nil {
			log.Fatal(err)

		}
		books = append(books, prd)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, bk := range books {
		fmt.Printf("%d - %s, %d, %d, %d, %s\n", bk.ID, bk.Bookname, bk.Pagenumber, bk.Stock, bk.Price, bk.Author)
	}
}

// func GetBookNameByID(id int) {
// 	var book string
// 	err := db.QueryRow("SELECT bookname FROM books WHERE id=$1", id).Scan(&book)
// 	switch {
// 	case err == sql.ErrNoRows:
// 		log.Printf("No book with that ID.")
// 	case err != nil:
// 		log.Fatal(err)
// 	default:
// 		fmt.Printf("Name of book  is %s\n", book)
// 	}

// }
func BuyBook(amount int, nameofbook string) {
	_, err = db.Exec("UPDATE books SET stock=stock-$1 WHERE bookname=$2", amount, nameofbook)
	if err != nil {
		log.Fatal(err)

	}
}

func DeleteBook(nameofbook string) {
	_, err = db.Exec("DELETE FROM books WHERE bookname=$1", nameofbook)
	if err != nil {
		log.Fatal(err)

	}

}

func GetBooksByID(id int) {

	var bookname string
	var pagenumber int
	var stock int
	var price int
	var author string
	err := db.QueryRow("SELECT * FROM books WHERE id=$1", id).Scan(&id, &bookname, &pagenumber, &stock, &price, &author)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No book with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Details of book: \n ID:  %d\n Bookname: %s \n Pagenumber: %d \n Stock: %d \n Price: %d \n Author: %s \n", id, bookname, pagenumber, stock, price, author)
	}
}
func GetBooksByName(bookname string) {

	var id int
	var pagenumber int
	var stock int
	var price int
	var author string
	err := db.QueryRow("SELECT * FROM books WHERE bookname=$1", bookname).Scan(&id, &bookname, &pagenumber, &stock, &price, &author)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No book with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Details of book: \n ID:  %d\n Bookname: %s \n Pagenumber: %d \n Stock: %d \n Price: %d \n Author: %s \n", id, bookname, pagenumber, stock, price, author)
	}
}
