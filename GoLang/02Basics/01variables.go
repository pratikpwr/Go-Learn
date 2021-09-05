package main

import "fmt"

func main() {

	// variable can be declared Multicase format
	// myName or MyName

	// 1 way
	var greetingText string       // variable declaration
	greetingText = "Hello World!" // variable initialization

	// 2 way
	// var variableName datatype = value
	var greets string = "Welcome to Go."

	// 3 way
	// using : before =
	greetings := "Have a Fun!"

	fmt.Println(greetingText)
	fmt.Println(greets)
	fmt.Println(greetings)
}
