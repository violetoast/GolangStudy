package main

import "fmt"

func main() {
	var myMap map[string]string
	if(myMap==nil) {
		fmt.Println("nil map")
	}
	myMap = make(map[string]string,10)
	myMap["one"]="java"
	myMap["two"]="c++"
	myMap["three"]="python"
	fmt.Println(myMap)

	//second 
	myMap2 := make(map[int]string)
	myMap2[1]="java"
	myMap2[2]="c++"
	myMap2[3]="python"
	fmt.Println(myMap2)

	//third
	myMap3 := map[string]string {
		"one" : "php",
		"two" : "c++",
		"three" : "python",
	}
	fmt.Println(myMap3)
}