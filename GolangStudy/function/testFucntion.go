package main

import "fmt"

func doublePrint(a int, b string) int {
	fmt.Println(a);
	fmt.Println(b);
	return 0
}

func doubleReturn(a int, b string) (int, int) {
	fmt.Println(a);
	fmt.Println(b);
	return 666,777
}

func doubleReturnPlus(a int, b string) (r1 int, r2 int) {
	fmt.Println(a);
	fmt.Println(b);
	r1,r2 = 100, 1000
	return
}

func doubleReturnPlusP(a int, b string) (r1, r2 int) {
	fmt.Println(a);
	fmt.Println(b);
	r1,r2 = 100, 1001
	return
}
func main() {
	r1, r2 := doubleReturnPlusP(1,"11")
	fmt.Println(r1, r2)
}