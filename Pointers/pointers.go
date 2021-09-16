package main

import "fmt"

func main () {
	age := 32;
	fmt.Println(age)

	myage := &age
	fmt.Println(*myage)

	// myage = 56
	*myage = 56
	fmt.Println(age)
	fmt.Println(*myage)
	fmt.Println("Same value")
	fmt.Println("Same values")
}