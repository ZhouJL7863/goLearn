package main

import (
	"fmt"
)

//这是简单题，并没有什么技巧
func removeDuplicates(nums []int) int {
	left, right, size := 0, 1, len(nums)
	for ; right < size; right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left], nums[right] = nums[right], nums[left]

	}
	return left + 1
}

func main() {
	nums := []int{1, 1, 2}
	ret := removeDuplicates(nums)
	fmt.Println(ret)
}
