package main

import "fmt"

func twoSum(nums []int, target int) []int {

	var ans []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				total := nums[i] + nums[j]
				if total == target {
					ans = append(ans, i, j)
					return ans
				}
			}
		}
	}
	return ans
}

func main() {

	arr := []int{4, 6, 8, 2}
	ans := twoSum(arr, 6)
	fmt.Println(ans)
}
