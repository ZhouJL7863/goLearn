package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 10, 25, 1}
	aim := 1000
	res := coins5(arr, aim)
	fmt.Println(res)
}

// func coins1(arr []int, aim int) int {
// 	if arr == nil || len(arr) == 0 || aim < 0 {
// 		return 0

// 	}
// 	return process1(arr, 0, aim)
// }

// func process1(arr []int, index, aim int) int {
// 	var res int = 0
// 	if index == len(arr) {
// 		if aim == 0 {
// 			res = 1
// 		} else {
// 			res = 0
// 		}
// 	} else {
// 		for i := 0; i*arr[index] <= aim; i++ {
// 			res += process(arr, index+1, aim-arr[index]*i)
// 		}
// 	}
// 	return res

// }

// //这是暴力解法，复杂度跟arr中钱的面值有关，最差情况为O(aim^N)

// func coins2(arr []int, aim int) {
// 	if arr == nil || len(arr) == 0 || aim < 0 {
// 		return 0
// 	}
// 	m := len(arr)
// 	var dp [][]int
// 	for i := 0; i < m; i++ {
// 		temp := make([]int, aim+1)
// 		dp = append(dp, temp)
// 	}
// 	return process2(arr, 0, aim, dp)
// }

// //新建一个map,可以重复利用计算值
// func process2(arr []int, index, aim int, dp [][]int) {
// 	var res int
// 	if index == len(arr) {
// 		if aim == 0 {
// 			res = 1
// 		} else {
// 			res = 0
// 		}
// 	} else {
// 		var mapValue int
// 		for i := 0; arr[index]*i <= aim; i++ {
// 			mapValue = dp[index+1][aim-arr[index]*i]
// 			if mapValue != 0 {
// 				if mapValue == -1 {
// 					res += 0
// 				} else {
// 					res += mapValue
// 				}
// 			} else {
// 				res += process2(arr, index+1, aim-arr[index]*i, dp)
// 			}
// 		}
// 	}
// 	if res == 0 {
// 		dp[index][aim] = -1
// 	} else {
// 		dp[index][aim] = res
// 	}
// 	return res
// }

// func coins3(arr []int, aim int) int {
// 	if arr == nil || len(arr) == 0 || aim < 0 {
// 		return 0
// 	}
// 	m := len(arr)
// 	var dp [][]int
// 	for i := 0; i < m; i++ {
// 		temp := make([]int, aim+1)
// 		dp = append(dp, temp)
// 	}
// 	for i := 0; i < m; i++ {
// 		dp[i][0] = 1
// 	}
// 	for j := 1; arr[0]*j <= aim; j++ {
// 		dp[0][arr[0]*j] = 1
// 	}

// 	var num int
// 	for i := 1; i < m; i++ {
// 		for j := 1; j <= aim; j++ {
// 			num = 0
// 			for k := 0; j-arr[i]*k >= 0; k++ {
// 				num += dp[i-1][j-arr[i]*k]
// 			}
// 			dp[i][j] = num
// 		}
// 	}
// 	return dp[m-1][aim]
// }

// //方法四真的是让人拍案叫绝！！！
// func coins4(arr []int, aim int) int {
// 	if len(arr) == 0 || aim < 0 {
// 		return 0
// 	}
// 	m := len(arr)
// 	var dp [][]int
// 	for i := 0; i < m; i++ {
// 		temp := make([]int, aim+1)
// 		dp = append(dp, temp)
// 	}
// 	for i := 0; i < m; i++ {
// 		dp[i][0] = 1
// 	}
// 	for j := 1; arr[0][j] <= aim; j++ {
// 		dp[0][arr[0]*j] = 1
// 	}
// 	for i := 1; i < m; i++ {
// 		for j := 1; j <= aim; j++ {
// 			dp[i][j] = dp[i-1][j]
// 			if j-arr[i] > 0 {
// 				dp[i][j] += dp[i][j-arr[i]]
// 			}
// 		}
// 	}
// 	return dp[m-1][aim]
// }

func coins5(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
	}

	dp := make([]int, aim+1)
	for j := 0; arr[0]*j <= aim; j++ {
		dp[arr[0]*j] = 1
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j <= aim; j++ {
			if j-arr[i] >= 0 {
				dp[j] += dp[j-arr[i]]
			} else {
				dp[j] += 0
			}
		}
	}
	return dp[aim]
}
