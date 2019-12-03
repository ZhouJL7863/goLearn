package main

import (
	"fmt"
	"sort"
	//"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			s := nums[i] + nums[start] + nums[end]
			if abs(target-s) < abs(target-ans) {
				ans = s
			}
			if s > target {
				end--
			} else if s < target {
				start++
			} else {
				return ans
			}
		}
	}
	return ans
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	ans := threeSumClosest(nums, target)
	fmt.Println(ans)
}
