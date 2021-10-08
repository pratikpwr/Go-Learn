package makePackage

import "fmt"

func makeFunc() {
	// make creates variables and reserve space in advance for them
	hobbies := make([]string, 2, 10)
	hobbies[0] = "Coding"
	hobbies[1] = "Anime"
	// above length will be 2 but capacity will be 10
	fmt.Println(hobbies)

	// in general case when we use append it creates new array
	// but when we use 'make()' to create vars  uses them
	hobbies = append(hobbies, "Gaming")
	fmt.Println(hobbies)
}
