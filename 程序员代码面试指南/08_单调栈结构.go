//题目描述，给定一个数组，找出每一个i位置左边和右边离i位置最近且值比arr[i]小的位置，返回所有位置的相应的信息

package main

import (
	"fmt"
	"stack"
)

func getNearLessNoRepeat(arr []int) [][2]int {
	n := len(arr)
	res := [][2]int{}
	var popIndex int
	var leftLessIndex int
	for i := 0; i < n; i++ {
		temp := [2]int{}
		res = append(res, temp)
	}
	var NewStack stack.Stack
	for i := 0; i < n; i++ {
		for NewStack.Len() != 0 && arr[NewStack.Top()] > arr[i] {
			popIndex = NewStack.Pop()
			if len(NewStack) == 0 {
				leftLessIndex = -1
			} else {
				leftLessIndex = NewStack.Top()
			}
			res[popIndex][0] = leftLessIndex
			res[popIndex][1] = i
		}
		NewStack.Push(i)
	}

	for NewStack.Len() != 0 {
		popIndex = NewStack.Pop()
		if NewStack.Len() == 0 {
			leftLessIndex = -1
		} else {
			leftLessIndex = NewStack.Top()
		}
		res[popIndex][0] = leftLessIndex
		res[popIndex][1] = -1
	}
	return res

}

func main() {
	arr := []int{3, 4, 1, 5, 6, 2, 7}
	res := getNearLessNoRepeat(arr)
	for _, k := range res {
		fmt.Println(k)
	}
}
