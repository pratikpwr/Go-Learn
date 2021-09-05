package main

import "fmt"

func main() {

	// variable can be declared Multicase format
	// myName - when var is used only in same package
	// MyName - when var is needed in another package

	// 1 way
	var greetingText string       // variable declaration
	greetingText = "Hello World!" // variable initialization

	// 2 way
	// var variableName datatype = value
	var greets string = "Welcome to Go."

	// 3 way
	// instantly initializing variable
	var gree = "Bye!"

	// 4 way
	// using : before =
	greetings := "Have a Fun!"

	fmt.Println(greetingText)
	fmt.Println(greets)
	fmt.Println(greetings)
	fmt.Println(gree)
}
