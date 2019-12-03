package main

import (
	"fmt"
)

//提交既然超时了
func process(arr []int, L, R int) int {
	if L == R {
		return arr[L-1] * arr[L] * arr[R+1]
	}
	var max int
	max = maxNum(arr[L-1]*arr[L]*arr[R+1]+process(arr, L+1, R), arr[L-1]*arr[R]*arr[R+1]+process(arr, L, R-1))
	for i := L + 1; i < R; i++ {
		max = maxNum(max, arr[L-1]*arr[i]*arr[R+1]+process(arr, L, i-1)+process(arr, i+1, R))

	}
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxCoins1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	n := len(arr)
	help := make([]int, n+2)
	help[0] = 1
	help[n+1] = 1
	for i := 0; i < len(arr); i++ {
		help[i+1] = arr[i]
	}
	return process(help, 1, n)
}

func main() {
	arr := []int{3, 2, 5}
	res := maxCoins2(arr)
	fmt.Println(res)
}

func maxCoins2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	n := len(arr)
	help := make([]int, n+2)
	help[0] = 1
	help[n+1] = 1
	for i := 0; i < len(arr); i++ {
		help[i+1] = arr[i]
	}
	fmt.Println(help)

	//生成一个dp[n+2][n+2]的矩阵
	var dp [][]int
	for i := 0; i < n+2; i++ {
		temp := make([]int, n+2)
		dp = append(dp, temp)
	}

	for i := 1; i <= n; i++ {
		dp[i][i] = help[i-1] * help[i] * help[i+1]
		fmt.Println(dp[i][i])
	}

	for i := n; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			var finalL int
			finalL = help[i-1]*help[i]*help[j+1] + dp[i+1][j]
			finalR := help[i-1]*help[j]*help[j+1] + dp[i][j-1]
			dp[i][j] = maxNum(finalL, finalR)
			for L := i + 1; L < j; L++ {
				dp[i][j] = maxNum(dp[i][j], help[i-1]*help[L]*help[j+1]+dp[i][L-1]+dp[L+1][j])
			}
			fmt.Println(i, j, dp[i][j])
		}

	}
	return dp[1][n]
}
