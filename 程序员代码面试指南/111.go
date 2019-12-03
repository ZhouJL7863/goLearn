package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var num int
	num = numDecodings(string(s[1:]))
	if k, _ := strconv.Atoi(string(s[:2])); k < 27 {
		num += numDecodings(string(s[2:]))
	}
	if len(s) > 2 {
		return num
	} else {
		return num + 1
	}
}

func main() {
	s := "12"
	res := numDecodings(s)
	fmt.Println(res)
}
