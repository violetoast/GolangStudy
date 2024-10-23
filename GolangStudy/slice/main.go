package main

import "fmt"

func solute(array [10]int) int{
	fmt.Println("------------")
	for _, value :=range array {
		fmt.Println(value)
	}
	return 0
}
func change(array [10]int) int{
	array[1]=100
	return 0
}
func main() {
	var a [10]int
	solute(a)
	change(a)
	solute(a)
	
	b:=[10]int{1,2,3,4}
	solute(b)
	change(b)
	solute(b)
}