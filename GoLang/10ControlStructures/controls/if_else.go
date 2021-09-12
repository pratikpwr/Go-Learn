package ifelse

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Enter your age: ")

	// reader := bufio.NewReader(os.Stdin)
	// userAgeInput, _ := reader.ReadString('\n')
	// userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	// userAge, _ := strconv.ParseInt(userAgeInput, 0, 64)

	// var userAgeInput string
	// fmt.Scanln(&userAgeInput)

	userAge, err := getUserInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	// if, else if, else in GO
	// we can also use && and || operators
	if userAge >= 18 {
		fmt.Println("You can access the content!")
	} else if userAge <= 8 {
		fmt.Println("You are baby, watch Cartoon!")
	} else {
		fmt.Println("Visit us when you are 18 and above!")
	}
}

func getUserInput() (int64, error) {
	var userAgeInput string
	fmt.Scanln(&userAgeInput)

	userIntInput, err := strconv.ParseInt(userAgeInput, 0, 64)

	return userIntInput, err

}
