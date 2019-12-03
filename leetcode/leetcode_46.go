//回溯算法（深度优先搜索+状态重置）
package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	n := len(nums)
	//用来表示一组可能的解
	vector := make([]int, n)
	//taken[i] == true 表示nums[i]已经出现在了vector中
	taken := make([]bool, n)
	fmt.Println(taken)
	ans := [][]int{}
	makePermutation(0, n, nums, vector, taken, &ans)
	return ans
}
func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if !taken[i] {
			taken[i] = true
			vector[cur] = nums[i]
			makePermutation(cur+1, n, nums, vector, taken, ans)
			taken[i] = false
		}
	}
}
func main() {
	nums := []int{1, 2, 3, 4}
	res := permute(nums)
	fmt.Println(len(res))
	for i := range res {
		fmt.Println(res[i])
	}

}
