package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	n := len(nums)
	info := false
	for j := n - 1; j >= 1; j-- {
		if nums[j] > nums[j-1] {
			for i := n - 1; i >= j; i-- {
				if nums[i] > nums[j-1] {
					nums[i], nums[j-1] = nums[j-1], nums[i]
					sort.Ints(nums[j:])
					fmt.Println(nums)
					break
				}

			}
			info = true

		}

		if info == true {
			break
		}
		if j == 1 {
			sort.Ints(nums)
			break
		}
		fmt.Println(nums[j])

	}
	//fmt.Println(nums)
}
func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
