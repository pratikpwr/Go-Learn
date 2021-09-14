package main

import "fmt"

// Correct Ans
func main() {

	input := "ab-cd" //"a-bC-dEf-ghIj"
	fmt.Println(input)
	str := reverseOnlyLetters(input)
	fmt.Println(str)

}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func reverseOnlyLetters(s string) (ans string) {

	// if not ltter store in map index and value
	// then remove it
	// reverse the stting formed
	// add !letter at position they were

	var odds map[int]string = map[int]string{}
	var newStr string
	for i := 0; i < len(s); i++ {
		isLetter := s[i] >= 97 && s[i] <= 122 || s[i] >= 65 && s[i] <= 90
		if !isLetter {
			odds[i] = string(s[i])
		} else {
			newStr = newStr + string(s[i])
		}
	}

	// fmt.Println(newStr)
	newStr = Reverse(newStr)
	// fmt.Println(newStr)

	var cnt int = 0
	for j := 0; j < len(s); j++ {

		if odds[j] != "" {
			ans = ans + string(odds[j])
		} else {
			ans = ans + string(newStr[cnt])
			cnt++
		}

	}
	return
}
