package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	} else if n < 0 {
		n++
		return 1 / x * myPow(x, n)
	} else {
		n--
		return x * myPow(x, n)
	}

}

func main() {
	x := 2.0
	n := -2
	ret := myPow(x, n)
	fmt.Println(ret)
}

//暴力解法内存溢出了（考虑清楚为什么会内存溢出 stack overflow)
//时间复杂度为O(n),空间复杂度为O(1).
