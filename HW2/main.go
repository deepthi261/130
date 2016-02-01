//Create a program that shows the type of some variable(use fmt.Printf)

package main

import "fmt"

func main() {

	var myInt int
	var myFloat float64
	myString := "My name is deepthi"
	myChar := 'A'

	fmt.Printf("%T \n", myInt)
	fmt.Printf("%T \n", myFloat)
	fmt.Printf("%T \n", myString)
	fmt.Printf("%T \n", myChar)

}


