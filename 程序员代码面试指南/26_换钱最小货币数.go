package main

import (
	"fmt"
)

func minCoins(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return -1
	}
	return process(arr, 0, aim)
}

func process(arr []int, i, rest int) int {
	//此时已经没有货币可以考虑了，如果此时剩余的钱为0，返回0张
	//如果此时剩余的钱不是0，返回-1
	if i == len(arr) {
		if rest == 0 {
			return 0
		}
		return -1
	}
	//最小张数，初始化为-1,因为还没有找到最优解
	var res int = -1
	for k := 0; k*arr[i] <= rest; k++ {
		//使用了k张arr[i],剩下的钱为rest-k*arr[i]
		//交给剩下的面值去考虑
		next := process(arr, i+1, rest-k*arr[i])
		if next != -1 {
			if res == -1 {
				res = next + k
			} else {
				res = minNum(res, next+k)
			}
		}
	}
	return res
}

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	arr := []int{5, 3, 2, 1}
	aim := 14
	res := minCoins(arr, aim)
	fmt.Println(res)
	res2 := minCoins2(arr, aim)
	fmt.Println(res2)
}

func minCoins2(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return -1
	}
	var N int = len(arr)
	var dp [][]int
	for i := 0; i <= N; i++ {
		temp := make([]int, aim+1)
		dp = append(dp, temp)
	}
	//除dp[N][0]之外都是-1
	for col := 1; col <= aim; col++ {
		dp[N][col] = -1
	}
	for i := N - 1; i >= 0; i-- {
		for rest := 0; rest <= aim; rest++ {
			dp[i][rest] = -1 //初始值先设置dp[i][rest]的值无效
			if dp[i+1][rest] != -1 {
				dp[i][rest] = dp[i+1][rest] //先设置成下面的值

			}
			if rest-arr[i] >= 0 && dp[i][rest-arr[i]] != -1 {
				if dp[i][rest] == -1 {
					dp[i][rest] = dp[i][rest-arr[i]] + 1
				} else {
					dp[i][rest] = minNum(dp[i][rest], dp[i][rest-arr[i]]+1)
				}
			}
		}
	}
	for i := 0; i <= N; i++ {
		fmt.Println(dp[i])
	}
	return dp[0][aim]

}
