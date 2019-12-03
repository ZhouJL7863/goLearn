package main

import (
	"fmt"
)

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 && n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	var temp int
	dp[0] = 0
	for j := 1; j <= n; j++ {
		dp[j] = j
	}
	for i := 1; i <= m; i++ {
		temp = dp[0]
		dp[0] = i
		for j := 1; j <= n; j++ {
			var res int
			if j > 0 {
				res = dp[j]
			}
			if word1[i-1] == word2[j-1] {
				dp[j] = temp
			} else if dp[j-1] == dp[j] {
				dp[j] = temp + 1
			} else {
				dp[j] = minNum(dp[j], dp[j-1]) + 1
			}
			temp = res
		}
		fmt.Println(dp)
	}
	return dp[n]
}

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	word1 := "sea"
	word2 := "ate"
	res := minDistance(word1, word2)
	fmt.Println(res)
}
