package main

import (
	"fmt"
)

func getdp2(arr []int) []int {
	dp := make([]int, len(arr))
	ends := make([]int, len(arr))
	ends[0] = arr[0]
	dp[0] = 1
	var right int
	var l int
	var r int
	var m int
	for i := 1; i < len(arr); i++ {
		l = 0
		r = right
		for l <= r {
			m = (l + r) / 2
			if arr[i] > ends[m] {
				l = m + 1
			} else {
				r = m - 1

			}
		}
		right = maxNum(right, l)
		ends[l] = arr[i]
		dp[i] = l + 1
	}
	fmt.Println(dp)
	return dp
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateLIS(arr, dp []int) []int {
	var L int
	var index int
	for i := 0; i < len(dp); i++ {
		if dp[i] > L {
			L = dp[i]
			index = i
		}
	}
	lis := make([]int, L)
	L--
	lis[L] = arr[index]
	for i := index; i >= 0; i-- {
		if arr[i] > arr[index] && dp[i] == dp[index]-1 {
			L--
			lis[L] = arr[i]
			index = i
		}
	}
	return lis
}

func main() {
	arr := []int{2, 2}
	dp := getdp2(arr)
	res := generateLIS(arr, dp)
	for _, value := range res {
		fmt.Println(value)
	}
}
