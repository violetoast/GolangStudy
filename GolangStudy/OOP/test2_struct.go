package main

import "fmt"

type Book struct {
	auth string
	title string
}

func changeBook(book *Book) {
	book.title = "see u again"
}
func main() {
	var book1 Book
	book1.title = "what can I say"
	book1.auth = "man"
	fmt.Println(book1)
	changeBook(&book1)
	fmt.Println(book1)

}