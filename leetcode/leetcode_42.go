package main

import (
	"fmt"
	"stack"
)

//接雨水，利用栈的思想来解决此题

func trap(height []int) int {
	n := len(height)
	var newStack stack.Stack
	var res int = 0
	for i := 0; i < n; i++ {
		for newStack.Len() != 0 && height[newStack.Peek()] < height[i] {
			Index := newStack.Pop()
			if newStack.Len() != 0 {
				temp := minNum(height[newStack.Peek()], height[i]) - height[Index]
				res = res + temp*(i-newStack.Peek()-1)
				fmt.Println(temp)
			}
		}

		newStack.Push(i)
	}
	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//这种方法不行，只能求取离i最近比它大的值
