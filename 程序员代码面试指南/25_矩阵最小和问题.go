package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}
	res := minPathSum(arr)
	fmt.Println(res)
}

func minPathSum(m [][]int) int {
	if m == nil || len(m) == 0 || len(m[0]) == 0 {
		return 0
	}

	more := MaxNum(len(m), len(m[0]))
	less := minNum(len(m), len(m[0]))
	var rowmore bool = more == len(m)
	arr := make([]int, less)
	arr[0] = m[0][0]
	for i := 1; i < less; i++ {
		if rowmore {
			arr[i] = arr[i-1] + m[0][i]
		} else {
			arr[i] = arr[i-1] + m[i][0]
		}
	}
	for i := 1; i < more; i++ {
		if rowmore {
			arr[0] = arr[0] + m[i][0]
		} else {
			arr[0] = arr[0] + m[0][i]
		}
		for j := 1; j < less; j++ {
			if rowmore {
				arr[j] = minNum(arr[j-1], arr[j]) + m[i][j]
			} else {
				arr[j] = minNum(arr[j-1], arr[j]) + m[j][i]
			}
		}
	}
	return arr[less-1]
}
func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func MaxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
