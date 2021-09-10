package main

import (
	"fmt"
)

func main() {

	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthdate := getUserData("Enter your birthdate (DD/MM/YYYY): ")

	// var newUser User = CreateUser(firstName, lastName, birthdate)
	var newUserUsingPointer User = *CreateUserUsingPointer(firstName, lastName, birthdate)

	// fmt.Println(firstName, lastName, birthdate, created)
	// fmt.Println(newUser)
	// fmt.Println(newUserUsingPointer)
	// outputUerData(&newUserUsingPointer)
	newUserUsingPointer.outputData()
}

func getUserData(promptText string) (data string) {
	fmt.Print(promptText)
	fmt.Scanln(&data)
	return
}
