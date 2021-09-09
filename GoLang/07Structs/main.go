package main

import (
	"fmt"
	"time"
)

// Struct in Go
// by this way struct can be defined
type User struct {
	firstname string
	lastname  string
	birthdate string
	createdAt time.Time
}

func main() {

	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthdate := getUserData("Enter your birthdate (DD/MM/YYYY): ")
	created := time.Now()

	var newUser User = User{
		firstname: firstName,
		lastname:  lastName,
		birthdate: birthdate,
		createdAt: created,
	}

	// fmt.Println(firstName, lastName, birthdate, created)
	fmt.Println(newUser)
}

func getUserData(promptText string) (data string) {
	fmt.Print(promptText)
	fmt.Scanln(&data)
	return
}
