package main

import (
	"fmt"
)

//题目描述:输入关于括号的字符串，问你是否是有效的括号， 即"{}[]" or "{[]}"都是有效的字符串
func isVaild(s string) bool {
	l := len(s)
	stack := make([]byte, l)
	top := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			stack[top] = c + 1
			top++
		} else if c == '[' || c == '{' {
			stack[top] = c + 2
			top++
		} else {
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
	}
	return top == 0
}
func main() {
	s := "{[]}"
	fmt.Println(len(s))
	ans := isVaild(s)
	fmt.Println("s", ans)
}
