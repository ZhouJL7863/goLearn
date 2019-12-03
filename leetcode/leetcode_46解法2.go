package main

import "fmt"
import "sort"

func permute(nums []int) [][]int {
	//用邹博讲的最小最大数来计算
	sort.Ints(nums)
	ret := [][]int{}
	temp2 := []int{}
	for i := range nums {
		temp2 = append(temp2, nums[i])
	}
	ret = append(ret, temp2)
	//fmt.Println(ret)
	j := len(nums) - 1
	temp := []int{}
	for j >= 1 {
		if nums[j] <= nums[j-1] {
			j--
			continue
		} else {
			for i := len(nums) - 1; i > j-1; i-- {
				if nums[i] > nums[j-1] {
					nums[j-1], nums[i] = nums[i], nums[j-1]
					sort.Ints(nums[j:])
					//fmt.Println(nums)
					for k := range nums {
						temp = append(temp, nums[k])
					}
					ret = append(ret, temp)
					temp = []int{}

					//fmt.Println(ret)
					break
				}

			}
			j = len(nums) - 1
		}
	}
	return ret
}
func main() {
	nums := []int{1, 2, 3, 4}
	ret := permute(nums)
	fmt.Println(ret)
}
