package main

import (
	"fmt"
	"strconv"
	"strings"

	"bmi/myconst"
)

func bmi() {

	fmt.Printf("%v %v %v\n", myconst.Separator, myconst.AppName, myconst.Separator)
	fmt.Print(myconst.Wprompt)
	// read the cmd line argu untill line terminated
	wInput, _ := reader.ReadString('\n') // this returns two values

	fmt.Print(myconst.Hprompt)
	hInput, _ := reader.ReadString('\n')

	// remove the line terminator
	wInput = strings.Replace(wInput, "\n", "", -1)
	hInput = strings.Replace(hInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(wInput, 64) // this returns two values
	height, _ := strconv.ParseFloat(hInput, 64)

	myBMI := weight / (height * height)

	fmt.Printf("Your BMI is %.2f", myBMI)
}
