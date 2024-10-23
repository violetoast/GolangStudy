package main 

import "fmt"

func main() {
	var array []int=make([]int,3,5)
	fmt.Printf("len=%d,cap=%d,array=%v\n",len(array),cap(array),array)
	array=append(array,1)
	fmt.Printf("len=%d,cap=%d,array=%v\n",len(array),cap(array),array)
	array=append(array,2)
	fmt.Printf("len=%d,cap=%d,array=%v\n",len(array),cap(array),array)
	array=append(array,3)
	fmt.Printf("len=%d,cap=%d,array=%v\n",len(array),cap(array),array)

}