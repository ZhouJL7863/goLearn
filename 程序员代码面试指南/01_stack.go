package main

import (
	"fmt"
)

//对应书 page.2 ，参考流程图1-1
//值类型包括：结构体，字符串，数组
//引用类型：切片，管道，map,指针
type Mystack struct {
	StackData []int
	StackMin  []int
}

func main() {
	arr := []int{3, 4, 5, 1, 2, 1}
	var stack Mystack //var this *Mystack 是错的，不能操作没用合法指向的内存
	this := &stack
	for i := 0; i < len(arr); i++ {
		newNum := arr[i]
		this.Push(newNum)
	}
	for i := 0; i < len(arr)-1; i++ {
		value := this.Pop()
		fmt.Println(value)
	}
	fmt.Println(this.StackData)
	fmt.Println(this.StackMin)
}

//实现了弹出操作
func (this *Mystack) Pop() (value int) {
	if len(this.StackData) == 0 {
		fmt.Println("Your stack is empty.")
		return 0
	} else {
		value = this.StackData[len(this.StackData)-1]
		if value == this.GetMin() {
			this.StackMin = this.StackMin[:len(this.StackMin)-1]
		}
		this.StackData = this.StackData[:len(this.StackData)-1]
	}
	return value
}

//实现了压入操作
func (this *Mystack) Push(newNum int) {
	if len(this.StackMin) == 0 {
		this.StackMin = append(this.StackMin, newNum)
	} else if newNum <= this.GetMin() {
		this.StackMin = append(this.StackMin, newNum)
	}
	this.StackData = append(this.StackData, newNum)
}

//实现了求栈的最小值
func (this *Mystack) GetMin() (value int) {
	if len(this.StackMin) == 0 {
		fmt.Println("Your stack is empty")
	}
	return this.StackMin[len(this.StackMin)-1]
}
