package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&n1)
	fmt.Print("Please enter a smaller number: ")
	fmt.Scan(&n2)
	fmt.Println(n1, "%", n2, " = ", n1%n2)
}
