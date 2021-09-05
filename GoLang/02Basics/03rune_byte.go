package main

import "fmt"

func main() {
	// rune and byte save/prints value as Unicode
	// use string type conv

	//rune (supports unique characters and emojis) but prints its UNIcode value
	var myRune rune = '₹'
	fmt.Println(myRune)         // prints UNIcode value => 8377
	fmt.Println(string(myRune)) // => ₹

	//byte (just like char in cpp) but prints its UNIcode value
	var myByte byte = 'a'
	fmt.Println(myByte)         // prints UNIcode value => 97
	fmt.Println(string(myByte)) // => a
}
