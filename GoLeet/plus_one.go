package main

import "fmt"

func main() {

	arr := []int{9, 9}
	ansArr := plusOne(arr)

	fmt.Print(ansArr)
}

func plus1(digits []int) []int {

	l := len(digits)

	if digits[l-1] < 9 {
		digits[l-1] = digits[l-1] + 1
		return digits
	}

	carry := 1
	for i := l - 1; i >= 0; i-- {
		if digits[i]+carry == 10 {
			digits[i] = (digits[i] + carry) % 10
			carry = 1
		} else {
			digits[i] = carry + digits[i]
			carry = 0
			break
		}
	}

	if carry == 1 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}

func plusOne(digits []int) []int {

	var ans []int
	length := len(digits)
	for i := 0; i < length; i++ {
		if i == length-1 {
			if digits[i] == 9 {

				if length == 1 {
					ans = append(ans, 1)
				} else {
					ans[i-1] = digits[i-1] + 1
				}
				ans = append(ans, 0)
			} else {
				ans = append(ans, digits[i]+1)
			}
		} else {
			ans = append(ans, digits[i])
		}
	}
	return ans
}
