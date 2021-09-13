package main

import "fmt"

func main() {
	arr := []int{4, 6, 8, 2, 6}
	ans := removeElement(arr, 6)
	fmt.Println(ans)
}

func removeElement(nums []int, val int) int {

	var ans []int
	for i, v := range nums {
		if v == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	fmt.Println(ans)
	return len(nums)
}
