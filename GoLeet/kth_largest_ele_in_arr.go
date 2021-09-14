package main

// failed for some test cases
// where int was int64

import (
	"fmt"
	"sort"
	"strconv"
)

func kthLargestNumber(nums []string, k int) string {

	var ans []int

	for _, val := range nums {
		conv, _ := strconv.ParseInt(val, 0, 64)
		ans = append(ans, int(conv))
	}
	// sort.Ints(ans)
	// sort.Reverse(ans)

	sort.Sort(sort.Reverse(sort.IntSlice(ans)))

	stringAns := fmt.Sprintf("%v", ans[k-1])

	return stringAns
}

func main() {
	arr := []string{"5", "7", "3", "10"}
	ans := kthLargestNumber(arr, 3)

	fmt.Println(ans)
}
