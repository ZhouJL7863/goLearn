package main

import (
	"fmt"
)

func longestVaildParentheses(s string) int {
	var maxans int = 0
	n := len(s)
	dp := make([]int, len(s))
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxans = maxNum(maxans, dp[i])
	}
	return maxans
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	s := "())(()))"
	res := longestVaildParentheses(s)
	fmt.Println(res)
}
