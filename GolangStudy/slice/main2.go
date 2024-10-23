package main 

import "fmt"

func solute(arrays []int) int{
	fmt.Println("------------")
	for _, value :=range arrays {
		fmt.Println(value)
	}
	return 0
}

func change(array []int) int{
	array[1]=100
	return 0
}

func main() {
	 a := []int{1,2,3,4}
	solute(a)
	change(a)
	solute(a)

}