package newPackage

import "fmt"

func newFunc() {

	// new allocates memeory to variables
	// it returns the address so store it in pointer var
	hobbies := new([]string)
	// so above hobbies is a pointer array

	fmt.Println(hobbies)
	fmt.Println(*hobbies)

	*hobbies = append(*hobbies, "Gaming")
	fmt.Println(*hobbies)
}
