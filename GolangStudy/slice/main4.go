package main 

import "fmt"

func main() {
	s := []int {1,2,3}
	s2 := s[0:2]
	fmt.Printf("s1=%v,s2=%v\n",s,s2)
	s2[0]=100;
	fmt.Printf("s1=%v,s2=%v\n",s,s2)
	var s3 []int = make([]int,3)
	copy(s3,s)
	fmt.Printf("s1=%v,s2=%v,s3=%v\n",s,s2,s3)
	s3[0]=1000
	fmt.Printf("s1=%v,s2=%v,s3=%v\n",s,s2,s3)

}