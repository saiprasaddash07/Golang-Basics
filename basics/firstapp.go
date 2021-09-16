package main

import (
	"fmt"
	"github.com/orgname/firstapp/greeting"
)

func main() {
	// var text string
	// text = "Hello World";

	// var text = "Hello World";
	// var text string = "Hello World";

	// text := "Hello World"
	// fmt.Println(text)
	fmt.Println(greeting.Text)

	x := 10
	var y int = x + 10
	fmt.Println(x)
	fmt.Println(y)
	y = 25
	fmt.Println(y)
	y++
	fmt.Println(y)

	var nums float64 = 5.6
	var newnums float64 = float64(x) / 2
	fmt.Println(nums)
	fmt.Println(newnums)

	var istrue bool = true
	fmt.Println(istrue)

	var r rune = '😂' // can store non ascii chars also
	fmt.Println(r)
	fmt.Println(string(r))

	var b byte = 'a' // can store ascii chars
	fmt.Println(b)

	firstName := "Saiprasad"
	lastName := "Dash"
	fmt.Println(firstName + " " + lastName)
	fmt.Println(firstName)
	full := firstName + lastName
	age := 20
	// fmt.Println("Hi I am " + full + " and I am " + age + " years old.");
	fmt.Printf("Hi I am %v and I am %v years old.", full, age)
	// full = fmt.Sprintln(firstName," ",lastName);
	// full = fmt.Sprintf("%v %v", firstName, lastName)

	// fmt.Println("9"+1); // error
	// fmt.Println(int("9")+1); // error bcoz int("hello") is not defined here
}
