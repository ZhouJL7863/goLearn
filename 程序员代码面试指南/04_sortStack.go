package main

import (
	"fmt"
	"stack"
)

var stack1 stack.Stack
var help stack.Stack

func sortStackBystack(stack1, help stack.Stack) {
	for stack1.Len() != 0 {
		cur, _ := stack1.Pop()
		top, _ := help.Top()
		value1, ok1 := cur.(int)
		value2, ok2 := top.(int)

		if ok1 && ok2 {
			for help.Len() != 0 && value2 < value1 {
				value, _ := help.Pop()
				stack1.Push(value)
			}
		}
		help.Push(cur)
	}
	for help.Len() != 0 {
		value, _ := help.Pop()
		stack1.Push(value)
	}

}
func main() {
	arr := []int{3, 4, 5, 1, 2}
	for i := 0; i < len(arr); i++ {
		stack1.Push(arr[i])
	}
	sortStackBystack(stack1, help)
	fmt.Println(help)
	fmt.Println(stack1)
}
