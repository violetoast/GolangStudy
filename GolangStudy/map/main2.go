package main

import "fmt"

func printMap (capitalMap map[string]string) {
	for key, value := range capitalMap {
		fmt.Println(key, value)
	}
}

func changeMap (capitalMap map[string]string) {
	capitalMap["Russia"] = "Moscow"
}
func main() {
	capitalMap := make(map[string]string)
	capitalMap["China"] = "Beijing"
	capitalMap["UK"] = "London"
	capitalMap["USA"] = "NewYork"

	for key, value := range capitalMap {
		fmt.Println(key, value)
	}

	fmt.Println("------------")

	delete(capitalMap,"UK")

	for key, value := range capitalMap {
		fmt.Println(key, value)
	}

	fmt.Println("------------")

	capitalMap["USA"]="DC"

	for key, value := range capitalMap {
		fmt.Println(key, value)
	}

	fmt.Println("111111111111111")
	printMap(capitalMap)
	fmt.Println("2222222222222222")
	changeMap(capitalMap)
	printMap(capitalMap)
}