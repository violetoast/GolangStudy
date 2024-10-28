package main

import "fmt"

type Human struct {
	Name string
	Sex string
}

func (this *Human) Print() {
	fmt.Println(this.Name)
	fmt.Println(this.Sex)
}

func (this *Human) Walk() {
	fmt.Println("Human walk....")
}

func (this *Human) Eat() {
	fmt.Println("Human eat....")
}

type Superman struct {
	Human
	Level int
}

func (this *Superman) Eat() {
	fmt.Println("Superman eat....")
}

func (this *Superman) Fly() {
	fmt.Println("Superman fly....")
}

func main() {
	h := Human{Name: "zhang3", Sex: "male"}
	h.Print()
	h.Walk()
	h.Eat()
	// v := Superman{Human{Name: "li4",Sex: "male"},88}
	v := Superman{Human:Human{Name: "li4",Sex: "male"},Level:88}
	v.Walk()
	v.Eat()
	v.Fly()
}