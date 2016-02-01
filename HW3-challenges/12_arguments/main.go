package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{5, 6, 7, 8}
	foo(aSlice...)
	foo()
}

func foo(numbers ...int) {
	fmt.Println(numbers)
}
