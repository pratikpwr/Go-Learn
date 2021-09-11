package main

import "fmt"

func main() {

	// maps in go
	// map[keytype]valuetype{}
	// maps are dynamic, any value can be used as key like int, struct
	websites := map[string]string{
		"Google":   "www.google.com/",
		"Facebook": "www.facebook.com/",
	}

	// print the whole map
	fmt.Println(websites)

	// prints value of the key
	fmt.Println(websites["Google"])

	// adding new item to map
	// use unique key
	websites["Instagram"] = "www.instagram.com/"
	fmt.Println(websites)

	// deleting element from array
	delete(websites, "Facebook")
	fmt.Println(websites)
}
