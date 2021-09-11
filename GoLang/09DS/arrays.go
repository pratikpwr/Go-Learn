package main

import (
	"fmt"
)

func mai() {

	// arrays in go
	prices := [4]int{23, 34, 12, 54}
	fmt.Println(prices)

	var days [7]string = [7]string{"Monday", "Tuesday", "Wednesday"}

	fmt.Println(days[0])
	days[3] = "Thursday"
	fmt.Println(days)

	// slices = provides the some part of lists
	// slices reference to the same array
	// so changes done in slices will effect array accordingly
	fastDays := days[1:3]
	fmt.Println(fastDays)

	// other ways of using slices
	fd := days[1:]
	fmt.Println(fd)

	fastD := days[:2]
	fmt.Println(fastD)

	// slices can be extended to right
	fmt.Println(len(fastD), cap(fastD)) // => 2, 7

	// append an array
	bigArray := append(days[0:4], "Friday")
	fmt.Println(bigArray)

	// dynamic arrays in go can be created by slices
	var months []string = []string{"Jan", "Feb", "Mar"}
	fmt.Println(months)
	months = append(months, "Apr")
	fmt.Println(months)

}
