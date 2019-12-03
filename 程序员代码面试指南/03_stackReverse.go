package main

import (
	"fmt"
)

//应该好好理解下递归的用法

type Stack struct {
	stack []int
}

func getAndRemoveLastElement(this *Stack) (ret int) {
	result := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.stack) == 0 {
		return result
	} else {
		var last int
		last = getAndRemoveLastElement(this)
		this.stack = append(this.stack, result)
		fmt.Println(last)
		return last
	}

}
func reverse(this *Stack) {
	if len(this.stack) == 0 {
		return
	}
	i := getAndRemoveLastElement(this)
	reverse(this)
	this.stack = append(this.stack, i)
}

func main() {
	var s Stack
	this := &s
	this.stack = []int{1, 2, 3}
	reverse(this)
	fmt.Println(this.stack)
}
