package main

import (
	"errors"
	"fmt"
)

// loops in go
// 1. for i:=0;i<count;i++{}
// 2. for condition {}
// 3. for index, value := ramge array {}

// we can use continue and break inside loops

func main() {
	userChoice, err := getUserChoice()
	if err != nil {
		fmt.Println(err)
		return
	}

	if userChoice == "1" {
		sumUptoX()
	} else if userChoice == "2" {
		factorialUptoX()
	} else if userChoice == "3" {
		sumAllNumbers()
	} else {
		sumNumbersInList()
	}

}

func sumUptoX() {
	num, err := getInputNo("Enter the no X: ")
	if err != nil {
		fmt.Println("Invalid number input.")
	}

	var sum int = 0
	for i := 0; i <= num; i++ {
		sum += i
	}

	fmt.Printf("S`um of numbers upto %v is %v", num, sum)
}

func factorialUptoX() {
	num, err := getInputNo("Enter the no X: ")
	if err != nil {
		fmt.Println("Invalid number input.")
	}

	var factorial int = 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}

	fmt.Printf("%v! : %v", num, factorial)
}

func sumAllNumbers() {

	isWantInput := true
	sum := 0

	// no while loop in GO
	// this way for loop can be used as while
	for isWantInput {
		num, err := getInputNo("Enter the number or press any alphabet to Exit: ")
		// if err != nil {
		// 	isWantInput = false
		// }
		// when error will be true we will set isWant input false so !(err!=nil)
		isWantInput = err == nil
		sum += num
	}

	fmt.Printf("Sum is %v", sum)
}

func sumNumbersInList() {

	size, err := getInputNo("Enter the size of list: ")
	if err != nil {
		fmt.Println("Invalid number input.")
	}

	var arr []int

	for i := 0; i < size; i++ {
		num, err := getInputNo("")
		if err != nil {
			fmt.Println("Invalid number input.")
			// to ignore incorrect input
			// continue
		}
		arr = append(arr, num)
	}

	var sum int = 0
	for j := 0; j < size; j++ {
		sum += arr[j]
	}
	fmt.Printf("Sum of numbers %v is %v", arr, sum)
}

func getInputNo(prompt string) (int, error) {
	fmt.Print(prompt)
	var numInput int

	_, err := fmt.Scanln(&numInput)
	if err != nil {
		return 0, err
	}

	return int(numInput), err
}

func getUserChoice() (string, error) {

	fmt.Println("Please make your choice!")
	fmt.Println("1) Add up all the numbers upto number X")
	fmt.Println("2) Build the factorial upto number X")
	fmt.Println("3) Sum up all the entered numbers")
	fmt.Println("4) Sum all the the numbers in list")
	fmt.Print("Enter your choice: ")
	var userInput string
	fmt.Scanln(&userInput)

	if userInput == "1" || userInput == "2" || userInput == "3" || userInput == "4" {
		return userInput, nil
	}

	err := errors.New("INVALID INPUT")
	return "", err
}
