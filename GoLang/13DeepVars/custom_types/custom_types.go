package custom_types

import "fmt"

type Person struct {
	name string
	age  int
}

// custom types
// type customNum int
// complex datatypes can be shortned with type
type PersonData map[string]Person

func customTypes() {

	// instead of using such long types we can use custom types
	var people map[string]Person = map[string]Person{
		"p1": {"Pratik", 21},
	}

	var newPeople PersonData = PersonData{
		"p2": {"Ganesh", 21},
	}

	fmt.Println(people)
	fmt.Println(newPeople)
}
