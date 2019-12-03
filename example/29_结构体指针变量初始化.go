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
	//顺序初始化，每个成员必须初始化,别忘了去地址
	var p1 *Student = &Student{1, "mile", 'm', 18, "bj"}
	fmt.Println("p1= ", *p1)
	//指定成员初始化，没有初始化的成员赋值为0
	s2 := &Student{name: "mike", addr: "bj"}
	fmt.Printf("the type is %T\n", s2)
	fmt.Println("s2= ", *s2)
}
