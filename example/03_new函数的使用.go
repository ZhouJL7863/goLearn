package main

import "fmt"

func main() {
	//a := 10//整型变量a
	var p *int
	//指向一个合法内存
	// p = &a
	//p是*int，指向int类型
	p = new(int) //动态分配空间，无需关心生命周期和垃圾回收
	*p = 666
	fmt.Println("*p= ", *p)
	q := new(int)
	*q = 777
	fmt.Println("*q = ", *q)
}
