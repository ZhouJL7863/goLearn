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
	s := Student{1, "mike", 'm', 18, "bj"}
	test01(s) //值传递，形参无法改变实参
	fmt.Println("main= ", s)
}
func test01(s Student) {
	s.id = 666
	fmt.Println("test01= ", s)

}
