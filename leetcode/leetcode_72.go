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
	var dp [][]int
	for i := 0; i <= m; i++ {
		temp := make([]int, n+1)
		dp = append(dp, temp)
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if dp[i-1][j] == dp[i][j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = minNum(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	for i := 0; i <= m; i++ {
		fmt.Println(dp[i])
	}
	return dp[m][n]
}

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	word1 := "horse"
	word2 := "ros"
	res := minDistance(word1, word2)
	fmt.Println(res)
}
