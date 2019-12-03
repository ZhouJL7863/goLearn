package main

import (
	"fmt"
)

//参考书page.5 图1-3
/*
	1.stackPush必须一次性的把所有的数据压入stackPop中
	2.stackPop 不为空，绝对不能向stackPop压入数据
*/
type Queue struct {
	stackPush []int //压入栈
	stackPop  []int //弹出栈
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var queue Queue
	this := &queue
	for i := 0; i < 5; i++ {
		newNum := arr[i]
		this.Add(newNum)
	}
	fmt.Println(this.stackPush)
	fmt.Println(this.stackPop)
}

func (this *Queue) PushToPop() {
	if len(this.stackPop) == 0 {
		if len(this.stackPush) != 0 {
			for i := len(this.stackPush) - 1; i >= 0; i-- {
				this.stackPop = append(this.stackPop, this.stackPush[i])
			}
			this.stackPush = this.stackPush[:0]
		}
	}
}

func (this *Queue) Add(newNum int) {
	this.stackPush = append(this.stackPush, newNum)
	this.PushToPop()

}

func (this *Queue) Poll() (value int) {
	if len(this.stackPop) == 0 {
		fmt.Println("Queue is empty!")
		this.PushToPop()
	}
	value = this.Peek()
	this.stackPop = this.stackPop[:len(this.stackPop)-1]
	return value
}

func (this *Queue) Peek() (peek int) {
	if len(this.stackPop) == 0 {
		fmt.Println("Queue is empty!")
		this.PushToPop()
	}
	return this.stackPop[len(this.stackPop)-1]
}
