package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaa")
}
func testb(x int) {
	//设置recover
	defer func() {
		//recover() //可以打印panic的错误信息
		//fmt.Println(recover())
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}() //调用此匿名函数
	//fmt.Println("bbbbbbbb")
	//显式调用panic函数，导致程序中断
	var a [10]int
	a[x] = 111 //当x为20的时候，导致数组越界，产生一个panic,导致程序崩溃
	//panic("this is a panic test")
}
func testc() {
	fmt.Println("ccccccc")
}
func main() {
	testa()
	testb(20)
	testc()
}
