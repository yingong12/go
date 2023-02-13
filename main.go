package main

import "fmt"

type Num int

func (n Num) print() {
	fmt.Println(n)
}

func (n *Num) pprint() {
	fmt.Println(*n)
}
func main() {
	var num Num
	defer num.print()
	defer num.pprint()
	defer func() {
		num.print()
	}()
	defer func() {
		num.pprint()
	}()
	num = 3
}
