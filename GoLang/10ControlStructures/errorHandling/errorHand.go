package errorhandling

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	// // errorhandling in GO
	// // errors are not thrown in dart they are reutrned as second argument
	userAge, err := getUserIn()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userAge)

}

func getUserIn() (int64, error) {
	var userAgeInput string
	fmt.Scanln(&userAgeInput)

	userIntInput, err := strconv.ParseInt(userAgeInput, 0, 64)

	var myError error
	if err != nil {
		// creating your own error
		myError = errors.New("invalid input")
	}

	return userIntInput, myError

}
