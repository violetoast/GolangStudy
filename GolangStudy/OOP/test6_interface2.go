package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)
	value, flag := arg.(string)
	if !flag {
		fmt.Println("its not string")
	} else {
		fmt.Println(value)
	}
}

type Book struct {
	auth string
}
func main() {
	book :=Book{"Golang"}
	myFunc(book)
	myFunc(1111)
	myFunc("nb")
}