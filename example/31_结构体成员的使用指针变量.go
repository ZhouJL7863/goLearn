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
	//1,指针有合法指向后，才操作成员。
	//先定义一个普通结构体变量
	var s Student
	//定义一个指针变量，保存s的地址
	var p1 *Student
	p1 = &s
	p1.id = 1
	(*p1).name = "mike"
	p1.age = 19
	p1.addr = "bj"
	fmt.Println("p1= ", p1)
	//2.通过new申请一个结构体
	p2 := new(Student)
	p2.id = 1
	p2.name = "mike"
	p2.sex = 'm'
	p2.addr = "bj"
	fmt.Println("p2=", p2)

}
