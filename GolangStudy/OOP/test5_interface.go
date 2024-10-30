package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor()
	GetType()
}

type Dog struct {
	color string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping....")
}

func (this *Cat) GetColor() {
	fmt.Println(this.color)
}

func (this *Cat) GetType() {
	fmt.Println("Cat")
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping....")
}

func (this *Dog) GetColor() {
	fmt.Println(this.color)
}

func (this *Dog) GetType() {
	fmt.Println("Dog")
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	animal.GetColor()
}
func main() {
	// var animal AnimalIF
	
	// animal = &Dog{"Green"}

	// animal.Sleep()

	// animal = &Cat{"Red"}

	// animal.GetColor()

	dog := Dog{"Green"}
	cat := Cat{"Red"}
	showAnimal(&dog)
	showAnimal(&cat)
}	