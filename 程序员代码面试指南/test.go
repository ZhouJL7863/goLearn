package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := nums[0]
	sum := 0
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
func main() {
	arr := []int{-2, 1, 3}
	//arr := [][]int{{0, 0}}
	res := maxSubArray(arr)
	fmt.Println(res)

}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
