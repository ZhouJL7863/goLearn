//三数之和，用双指针求解
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	//排序后，可以按规律查找
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		//i表示的时nums的索引
		//i>0是防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				//需要的L小了
				l++

			} else if s > 0 {
				//需要的r大了
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l, r = next(nums, l, r)
			}

		}
	}
	return res
}
func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
