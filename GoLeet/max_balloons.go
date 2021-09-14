package main

import "fmt"

// failed for some test cases
func main() {

	c := maxNumberOfBalloons("lloo")
	fmt.Println(c)

}

func maxNumberOfBalloons(text string) int {

	ballonCount := map[string]int{
		"b": 1,
		"a": 1,
		"l": 2,
		"o": 2,
		"n": 1,
	}

	textCount := map[string]int{
		"b": 0,
		"a": 0,
		"l": 0,
		"o": 0,
		"n": 0,
	}

	for i := 0; i < len(text); i++ {
		if string(text[i]) == "b" {
			textCount["b"]++
		}
		if string(text[i]) == "a" {
			textCount["a"]++
		}
		if string(text[i]) == "l" {
			textCount["l"]++
		}
		if string(text[i]) == "o" {
			textCount["o"]++
		}
		if string(text[i]) == "n" {
			textCount["n"]++
		}
	}
	// fmt.Println(ballonCount)
	// fmt.Println(textCount)

	cnt0 := 0
	for k, v := range ballonCount {
		if textCount[k]%v == 0 {
			cnt0++
		}
	}
	ans := 0
	if cnt0 == 5 {
		ans = textCount["l"] / ballonCount["l"]
	}

	return ans
}
