package main

import "fmt"

type Person struct {
	name string
	age  int
}

type PersonData map[string]Person

func main() {

	var people map[string]Person = map[string]Person{
		"p1": {"Pratik", 21},
	}

	var newPeople PersonData = PersonData{
		"p2": {"Ganesh", 21},
	}

	fmt.Println(people)
	fmt.Println(newPeople)
}
