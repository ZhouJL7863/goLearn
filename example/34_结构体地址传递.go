package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	s1 := Student{1, "mike", 'm', 18, "bj"}
	test01(&s1)
	fmt.Println("main= ", s1)
}
func test01(p *Student) {
	*p.id = 666 //地址传递，形参可以改实参
	fmt.Println("p= ", *p)
}
