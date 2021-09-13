package main

import "fmt"

// Steps to solve

// 1. no. of rows = no of times outer loop will run

// 2. Identify for every row no
// - how many columns are there
// - types of element in column

// 3. what do you need to print

// try to find formula relative to rows and col

func main() {
	// var num int
	// fmt.Scanln(&num)
	pat10(5)
}

func pat10(n int) {

}

func pat09(n int) {
	//*********
	// *******
	//  *****
	//   ***
	//    *

	for row := n; row > 0; row-- {

		for col := 0; col < n-row; col++ {
			fmt.Print(" ")
		}

		for col := row*2 - 1; col > 0; col-- {
			fmt.Print("*")
		}

		fmt.Println("")
	}
}

func pat08(n int) {
	//     *
	//    ***
	//   *****
	//  *******
	// *********
	for rows := 1; rows <= n; rows++ {

		for col := 0; col < n-rows; col++ {
			fmt.Print(" ")
		}
		for col := 1; col <= 2*rows-1; col++ {
			fmt.Print("*")
		}
		// it will work if we remove below for loop
		for col := 0; col < n-rows; col++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func pat07(n int) {
	// *****
	//  ****
	//   ***
	//    **
	//     *
	for rows := 0; rows < n; rows++ {

		for columns := 0; columns < rows; columns++ {
			fmt.Print(" ")
		}
		for columns := n - rows; columns > 0; columns-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pat06(n int) {
	// 	    *
	// 	   **
	//    ***
	//   ****
	//  *****
	for rows := 1; rows <= n; rows++ {
		for cols := 0; cols <= n-rows; cols++ {
			fmt.Print(" ")
		}
		for cols := 0; cols < rows; cols++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func pat5(n int) {
	// 1
	// 2 1
	// 3 2 1
	// 4 3 2 1
	// 3 2 1
	// 2 1
	// 1
	for rows := 0; rows < 2*n; rows++ {

		var noOfCol int
		if rows > n {
			noOfCol = 2*n - rows
		} else {
			noOfCol = rows
		}

		for col := noOfCol; col > 0; col-- {
			fmt.Printf("%v ", col)
		}
		fmt.Println()
	}
}

func pat05(n int) {
	// *
	// * *
	// * * *
	// * * * *
	// * * *
	// * *
	// *
	for row := 0; row < 2*n; row++ {

		var totalColsInRow int
		if row > n {
			totalColsInRow = (2*n - row)
		} else {
			totalColsInRow = row
		}

		for col := 0; col < totalColsInRow; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func pat01(n int) {
	// * * * *
	// * * * *
	// * * * *
	// * * * *
	for row := 1; row <= n; row++ {
		for col := 0; col < n; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func pat02(n int) {
	// *
	// * *
	// * * *
	// * * * *
	for row := 1; row <= n; row++ {
		// for every row run the column
		for col := 0; col < row; col++ {
			fmt.Print("* ")
		}
		// when one row is printed , add a new line
		fmt.Println()
	}
}

func pat03(n int) {
	// * * * *
	// * * *
	// * *
	// *
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func pat04(n int) {
	// 1
	// 1 2
	// 1 2 3
	// 1 2 3 4
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}
}
