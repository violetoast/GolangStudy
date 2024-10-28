package main 

import "fmt"

type Hero struct {
	Name string
	Ad int
	Level int 
}

func (this *Hero) Show() {
	fmt.Println(this.Name)
	fmt.Println(this.Ad)
	fmt.Println(this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name=newName
}

func main() {
	hero := Hero{Name: "zhang3",Ad: 100,Level: 1}
	hero.Show()
	hero.SetName("li4")
	hero.Show()
}