package main

import (
	"fmt"
)

//对应书page.3 参考图1-2
type MyStack struct {
	stackData []int
	stackMin  []int
}

func main() {
	arr := []int{3, 4, 5, 1, 2, 1}
	var stack MyStack
	this := &stack
	for i := 0; i < len(arr); i++ {
		newNum := arr[i]
		this.Push(newNum)
	}
	fmt.Println(this.stackData)
	fmt.Println(this.stackMin)
}

func (this *MyStack) Push(newNum int) {
	this.stackData = append(this.stackData, newNum)
	if len(this.stackMin) == 0 {
		this.stackMin = append(this.stackMin, newNum)
	} else if newNum <= this.GetMin() {
		this.stackMin = append(this.stackMin, newNum)
	} else if newNum > this.GetMin() { //else 后面不能接条件语句
		this.stackMin = append(this.stackMin, this.GetMin())
	}
}

func (this *MyStack) Pop() (value int) {
	if len(this.stackData) == 0 {
		return 0
	}
	value = this.stackData[len(this.stackData)-1]
	this.stackData = this.stackData[:len(this.stackData)-1]
	this.stackMin = this.stackMin[:len(this.stackData)-1]
	return value
}

func (this *MyStack) GetMin() (value int) {
	if len(this.stackMin) == 0 {
		return 0
	}
	return this.stackMin[len(this.stackMin)-1]
}

//方案一和方案二对比
/*
	方案一和方案二其实都是用stackMin保存着stackData每一步的最小值，共同点是所有操作的复杂度都为O(1),空间复杂度为O(n).
	区别是：方案一中stackMin压入时间稍省空间，但是弹出操作费时间；方案二中stackMin压入稍费时间，但是弹出操作稍省时间。
*/
