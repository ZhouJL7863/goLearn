package main

import (
	"fmt"
)

func getdp(str1, str2 string) [][]int {
	var dp [][]int
	//m := len(str1)
	//n := len(str2)
	for i := 0; i < len(str1); i++ {
		temp := make([]int, len(str2))
		dp = append(dp, temp)
	}
	if str1[0] == str2[0] {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	for i := 1; i < len(str1); i++ {
		if str1[i] == str2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1; j < len(str2); j++ {
		if str2[j] == str1[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < len(str1); i++ {
		for j := 1; j < len(str2); j++ {
			dp[i][j] = MaxNum(dp[i-1][j], dp[i][j-1])
			if str1[i] == str2[j] {
				dp[i][j] = MaxNum(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp
}

func main() {
	str1 := "1A2C3D4B56"
	str2 := "B1D23CA45B6A"
	res := lcse(str1, str2)
	fmt.Println(res)
}

func MaxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcse(str1, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	dp := getdp(str1, str2)
	m := len(str1) - 1
	n := len(str2) - 1
	res := make([]byte, dp[m][n])
	index := len(res) - 1
	for index >= 0 {
		if m > 0 && dp[m][n] == dp[m-1][n] {
			m--
		} else if n > 0 && dp[m][n] == dp[m][n-1] {
			n--
		} else {
			res[index] = str1[m]
			m--
			n--
			index--
		}
	}
	return string(res)

}
