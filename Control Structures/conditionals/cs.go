package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getuserage() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your age: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	age, error := strconv.ParseInt(userInput, 0, 64)
	return int(age), error
}

func main() {
	age,error := getuserage()
	if error == nil {
		if age > 20 {
			fmt.Println("Hello")
		} else if age > 12 {
			fmt.Println("Hellos")
		}
	} else {
		fmt.Println("Invalid input")
		return
	}
}
