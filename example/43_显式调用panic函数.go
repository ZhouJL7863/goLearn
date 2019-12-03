package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaa")
}
func testb() {
	//fmt.Println("bbbbbbbb")
	//显式调用panic函数，导致程序中断
	panic("this is a panic test")
}
func testc() {
	fmt.Println("ccccccc")
}
func main() {
	testa()
	testb()
	testc()
}
