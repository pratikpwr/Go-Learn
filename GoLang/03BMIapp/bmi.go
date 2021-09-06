package main

import (
	"fmt"

	"bmi/myconst"
)

func main() {

	var weight float32
	var height float32
	var myBmi float32

	fmt.Printf("%v %v %v\n", myconst.Separator, myconst.AppName, myconst.Separator)
	fmt.Print(myconst.Wprompt)
	fmt.Scanln(&weight)
	fmt.Print(myconst.Hprompt)
	fmt.Scanln(&height)

	myBmi = weight / (height * height)

	fmt.Printf("Your BMI is %.2f", myBmi)

}
