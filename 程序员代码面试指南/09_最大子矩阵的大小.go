package main

import (
	"fmt"
	"stack"
)

func maxRecSize(arr [][]int) int {
	if arr == nil || len(arr) == 0 || len(arr[0]) == 0 {
		return 0
	}

	var maxArea int = 0
	height := make([]int, len(arr[0]))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 0 {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		maxArea = MaxNum(maxRecFromBottom(height), maxArea)
	}
	return maxArea
}

func maxRecFromBottom(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	maxArea := 0
	var newStack stack.Stack
	var k int
	for i := 0; i < len(height); i++ {
		for newStack.Len() != 0 && height[i] <= height[newStack.Peek()] {
			j := newStack.Pop()
			if newStack.Len() == 0 {
				k = -1
			} else {
				k = newStack.Peek()
			}
			curArea := (i - k - 1) * height[j]
			maxArea = MaxNum(maxArea, curArea)
		}
		newStack.Push(i)
	}
	for newStack.Len() != 0 {
		j := newStack.Pop()
		if newStack.Len() == 0 {
			k = -1
		} else {
			k = newStack.Peek()
		}
		curArea := (len(height) - k - 1) * height[j]
		maxArea = MaxNum(maxArea, curArea)
	}
	return maxArea
}

func MaxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := [][]int{
		{1, 0, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 0},
	}
	res := maxRecSize(arr)
	fmt.Println(res)
}

//这道题跟leetcode中求最大可积水面积相似
