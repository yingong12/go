package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [10]int{}
	b := []int{1, 2}
	func(a [10]int, b []int) {
		fmt.Println(unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]))
		a[0] = 1
		b[0] = 2
	}(a, b)
	fmt.Println(a, b, unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]))
}
