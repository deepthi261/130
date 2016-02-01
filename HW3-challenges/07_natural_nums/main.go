
package main

import "fmt"

func main() {
	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}
