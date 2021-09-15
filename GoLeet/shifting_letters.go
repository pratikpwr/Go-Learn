package main

import "fmt"

// Incomplete
func main() {
	str := "abc"
	arr := []int{3, 5, 9}
	ans := shiftingLetters(str, arr)
	fmt.Println(ans)
}

func shiftingLetters(s string, shifts []int) string {

	// take as seperate arr
	// perform inc operation on byte
	//

	var arr = []byte{}

	for i := 0; i < len(s); i++ {
		arr = append(arr, s[i])
	}

	cnt := 0
	for j := 0; j < cnt+1; j++ {
		for k := 0; k <= j; k++ {
			if j < 3 {
				arr[j] = arr[k] + byte(shifts[j])
			}
		}
		cnt++
	}

	fmt.Println(arr)
	return ""
}
