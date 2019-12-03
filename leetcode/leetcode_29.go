package main

import (
	"errors"
	"fmt"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		err := errors.New("除数不能为0")
		fmt.Println(err)
	}
	n := 0
	for {
		if dividend < divisor {
			break
		}
		dividend = dividend - divisor
		n++
	}
	return n
}
func main() {
	dividend := 30
	divisor := 3
	ret := divide(dividend, divisor)
	fmt.Println(ret)
}
