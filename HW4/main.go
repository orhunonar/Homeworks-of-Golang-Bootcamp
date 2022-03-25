package main

import (
	"flag"
	"homework4/models"
)

func main() {
	// book1 := models.BookInfo{
	// 	Bookname:   "Nutuk",
	// 	Pagenumber: 11,
	// 	Stock:      11,
	// 	Price:      11,
	// 	Author:     "Atatürk",
	// }
	// book2 := models.BookInfo{
	// 	Bookname:   "Silmarillion",
	// 	Pagenumber: 21,
	// 	Stock:      21,
	// 	Price:      21,
	// 	Author:     "Tolkien",
	// }
	// book3 := models.BookInfo{
	// 	Bookname:   "Olasılıksız",
	// 	Pagenumber: 30,
	// 	Stock:      40,
	// 	Price:      50,
	// 	Author:     "Adam Fawer",
	// }

	// models.InsertBook(book1)
	// models.InsertBook(book2)
	// models.InsertBook(book3)

	//models.GetBooks()

	// a := 3
	// models.BuyBook(a, "Nutuk")
	//models.GetBooks()

	// models.DeleteBook("Nutuk")
	// models.GetBooks()

	//models.GetBooksByID(4)

	listing := flag.Bool("list", false, "")
	searching := flag.Int("search", 0, "")
	deleting := flag.String("delete", "", "")
	purchasing := flag.String("buy", "", "")
	stocking := flag.Int("amount", 0, "")

	flag.Parse()

	if *listing {
		models.GetBooks()
	} else if *searching != 0 {
		models.GetBooksByID(*searching)
	} else if *deleting != "" {
		models.DeleteBook(*deleting)
	} else if *purchasing != "" && *stocking >= 0 {
		models.BuyBook(*stocking, *purchasing)
		models.GetBooksByName(*purchasing)
	}

}
