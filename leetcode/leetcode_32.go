package main

import (
	"fmt"
	"stack"
)

func longestVaildParentheses(s string) int {
	n := len(s)
	var newStack stack.Stack
	var max int = 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			newStack.Push(i)
		} else {
			newStack.Pop()
			if newStack.Len() == 0 {
				newStack.Push(i)
			} else {
				max = maxNum(max, i-newStack.Peek())
			}
		}
	}
	return max
}

func main() {
	s := ")()())"
	res := longestVaildParentheses(s)
	fmt.Println(res)
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
