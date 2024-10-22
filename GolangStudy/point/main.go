package main 

import "fmt"

func swap(a *int, b *int) {
	temp:=*a
	*a=*b
	*b=temp
}
func main() {
	a:=10
	b:=100;
	swap(&a,&b)
	fmt.Println(a, b)
}