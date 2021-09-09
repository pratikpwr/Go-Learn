package main

import "fmt"

func main() {
	age := 21

	fmt.Println(age)

	// just like pointers in cpp
	var myAge *int = &age

	fmt.Println(myAge)

	// using * at pointer variable we can access the value
	fmt.Println(*myAge)

	// we can also change value using *
	*myAge = 22
	fmt.Println(*myAge)
	fmt.Println(age)

	doubledAge := double(age)
	fmt.Println(doubledAge)

	// passing pointer as parameter
	doubleAge(myAge)
	fmt.Println(age)

}

func double(num int) int {
	return num * 2
}

// this functions takes pointer parameter
// changes done here will reflect in main var
func doubleAge(num *int) {
	*num = *num * 2
}
