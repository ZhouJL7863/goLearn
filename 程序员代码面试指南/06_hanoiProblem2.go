//这道题暂时保留到这
package main

import (
	"fmt"
	"stack"
	"strconv"
)

func hanoiProblem2(num int, left, mid, right string) int {
	Action := make([]string, 5)
	var MAX_VALUE int
	MAX_VALUE = 1<<32 - 1
	var Ls stack.Stack
	var Ms stack.Stack
	var Rs stack.Stack
	Ls.Push(MAX_VALUE)
	Ms.Push(MAX_VALUE)
	Rs.Push(MAX_VALUE)
	for i := num; i > 0; i-- {
		Ls.Push(i)
	}
	Action = append(Action, "NO")
	var step int
	step = 0
	for Rs.Len() != num+1 {
		step += fStackToStack(Action, "MTOL", "LTOM", Ls, Ms, left, mid)
		step += fStackToStack(Action, "LTOM", "MTOL", Ms, Ls, mid, left)
		step += fStackToStack(Action, "MTOR", "RTOM", Rs, Ms, right, mid)
		step += fStackToStack(Action, "RTOM", "MTOR", Ms, Rs, mid, right)
	}
	return step

}

func fStackToStack(Action []string, preNoAct, nowAct string, fstack, tstack stack.Stack, from, to string) int {
	if Action[0] != preNoAct && fstack.Peek() < tstack.Peek() {
		tstack.Push(fstack.Pop())
		fmt.Println("Move " + strconv.Itoa(tstack.Peek()) + " from " + from + " to " + to)
		Action[0] = nowAct
		return 1
	}
	return 0
}

func main() {
	num := 2
	left, right, mid := "left", "right", "mid"
	ret := hanoiProblem2(num, left, right, mid)
	fmt.Println(ret)
}
