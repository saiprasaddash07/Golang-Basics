package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthdate   string
	createdDate time.Time
}

func (user *User) outputDetails() {
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthdate)
}

func NewUser(fName string, lName string, bdate string) *User {
	created := time.Now()
	user := User{
		firstName:   fName,
		lastName:    lName,
		birthdate:   bdate,
		createdDate: created,
	}
	return &user
}

var reader = bufio.NewReader(os.Stdin)

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userInput, "\n", "", -1)
	return cleanedInput
}

// func outputUserDetails(user *User) {
// 	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthdate)
// }

func main() {
	var newUser *User

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	newUser = NewUser(firstName, lastName, birthdate)
	// outputUserDetails(newUser) 
	newUser.outputDetails()

	// newUser = User{firstName, lastName, birthdate, created}

	fmt.Println(newUser)
}
