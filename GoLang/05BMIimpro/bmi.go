package main

import (
	"bmi/myconst"
)

func main() {

	myconst.AppIntro()

	weight, height := getUserMatrics()

	myBMI := calcBMI(weight, height)

	printBMI(myBMI)
}

func calcBMI(weight float64, height float64) (bmi float64) {
	bmi = weight / (height * height)
	return
}
