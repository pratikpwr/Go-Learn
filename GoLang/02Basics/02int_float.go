package main

import "fmt"

func main() {

	var num int = 17
	num = num * 4

	fmt.Println(num)

	// Go is strictly typed
	// use explecite type conversion
	var myNum float64 = float64(num) / 3
	fmt.Println(myNum)

	// float64 vs float32

	var defaultFloat float64 = 5.934659328476587325987 // prints as it is
	var smallFloat float32 = 5.7289265659827487234     // prints with less precision

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)
}
