package main

import "fmt"

func test01() int {
	//函数被调用时，x才分配空间，才初始化为0
	var x int //没有初始化，值为0
	x++
	return x * x
}

func main() {
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
}
