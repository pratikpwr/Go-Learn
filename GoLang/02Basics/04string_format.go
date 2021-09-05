package main

import "fmt"

func main() {
	var firstName string = "Pratik"
	lastName := "Pawar"

	// fullName := firstName + " " + lastName
	fullName := fmt.Sprintf("%v %v", firstName, lastName)

	// fmt.Println() - prints with endline
	fmt.Println(fullName)

	age := 21

	// printing using format - like c lang
	fmt.Printf("Hi, My name is %v and I am %v years old.\n", fullName, age)

	// fmt.Sprintf() return formatted string but doesnt print
	myDesc := fmt.Sprintf("Hi, My name is %v and I am %v years old.", fullName, age)

	fmt.Println(myDesc)

	// %v  - variables
	// %+v - structs
	// %T - datatype of variable
	// %% - percentage symbol
	//%.2f - float with 2 decimal value
}
