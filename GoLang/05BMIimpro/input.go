package main

import (
	"bmi/myconst"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMatrics() (float64, float64) {
	var weight float64
	var height float64

	fmt.Print(myconst.Wprompt)
	fmt.Scanln(&weight)
	fmt.Print(myconst.Hprompt)
	fmt.Scanln(&height)
	return weight, height
}

func readUserMatrics() (weight float64, height float64) {
	weight = getUserInput(myconst.Wprompt)
	height = getUserInput(myconst.Hprompt)
	return
}

func getUserInput(prompt string) float64 {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	data, _ := strconv.ParseFloat(input, 64)
	return data
}
