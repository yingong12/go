package main

import "fmt"

type Mamal interface {
	Barkk()
}

type Dog struct {
	Name string
}

func (d *Dog) Bark() {
	d.Name = "ruirui"
	fmt.Println("woof")
}
func (d Dog) Barkk() {
}

func main() {
	dog := Dog{}
	var m Mamal
	m = dog
	//语法糖。 解析为Bark(*dog)
	m.Barkk()
}
