package main

import (
	"fmt"
	"stack"
)

func getNearLess(arr []int) [][2]int {
	n := len(arr)
	res := [][2]int{}
	for i := 0; i < n; i++ {
		temp := [2]int{}
		res = append(res, temp)
	}
	var NewStack []stack.Stack
	var LeftIndex int
	//var RightIndex int
	for i := 0; i < n; i++ {
		//弹出操作设置

		for len(NewStack) != 0 && arr[NewStack[len(NewStack)-1].Peek()] > arr[i] {
			stackTemp := NewStack[len(NewStack)-1]
			NewStack = NewStack[:len(NewStack)-1] //出栈
			if len(NewStack) == 0 {
				LeftIndex = -1
			} else {
				LeftIndex = NewStack[len(NewStack)-1].Peek()
			}
			for stackTemp.Len() != 0 {
				Index := stackTemp.Pop()
				res[Index][0] = LeftIndex
				res[Index][1] = i
			}
		}
		if len(NewStack) == 0 {
			var TempStack stack.Stack
			TempStack.Push(i)
			NewStack = append(NewStack, TempStack)
		} else {
			if arr[NewStack[len(NewStack)-1].Peek()] == arr[i] {
				NewStack[len(NewStack)-1].Push(i)
			} else {
				var TempStack stack.Stack
				TempStack.Push(i)
				NewStack = append(NewStack, TempStack)
			}
		}

	}
	for len(NewStack) != 0 {
		stackTemp := NewStack[len(NewStack)-1]
		NewStack = NewStack[:len(NewStack)-1] //出栈
		if len(NewStack) == 0 {
			LeftIndex = -1
		} else {
			LeftIndex = NewStack[len(NewStack)-1].Peek()
		}
		for stackTemp.Len() != 0 {
			Index := stackTemp.Pop()
			res[Index][0] = LeftIndex
			res[Index][1] = -1
		}
	}
	return res
}

func main() {
	arr := []int{3, 1, 3, 4, 3, 5, 3, 2, 2}
	res := getNearLess(arr)
	for _, k := range res {
		fmt.Println(k)
	}
}

//感受就是我们自己实现的这个栈太不智能啊
//接下来要是实现一个接口啊，要不然要死人的
//同时封装好一个队列
