package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		ans = maxNum(ans, sum)
	}
	return ans
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 3, -1, -5, 4, 3, -2}
	res := maxSubArray(arr)
	fmt.Println(res)
}
