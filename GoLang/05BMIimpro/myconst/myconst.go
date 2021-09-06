package myconst

import "fmt"

const appName = "BMI Calculator"
const separator = "---------"
const Wprompt = "Enter your weight in kg: "
const Hprompt = "Enter your height in m: "

func AppIntro() {
	fmt.Printf("%v %v %v\n", separator, appName, separator)
}
