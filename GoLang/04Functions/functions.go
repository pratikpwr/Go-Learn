package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var a int
	var b int
	// fmt.Print("Enter numbers to add: ")
	// fmt.Scanf("%d %d", &a, &b)

	a, b = genRandomNo()

	sum := add(a, b)
	fmt.Printf("Sum of %v+%v=%v", a, b, sum)
	sub := subtract(a, b)
	fmt.Printf("\nSubtraction of %v-%v=%v", a, b, sub)
}

// this is how function is declared in go
func add(num1 int, num2 int) int {
	return num1 + num2
}

// named return values
// no need to return value just call return
func subtract(num1 int, num2 int) (sub int) {
	sub = num1 - num2
	return
}

// functions in go can return one or more value
// specify the data types of return value in brackets
func genRandomNo() (randNo1 int, randNo2 int) {
	randNo1 = rand.Intn(10)
	randNo2 = rand.Intn(10)
	return randNo1, randNo2
}
