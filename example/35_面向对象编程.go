package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person //只有类型，没有字段
	id     int
	addr   string
}

func main() {
	//面向对象编程有三个性质：1封装，通过方法实现
	//2.继承，通过匿名字段实现
	//3.多态：通过接口实现
	var s1 Student = Student{Person{"yoyo", 'm', 12}, 20, "bj"}
	fmt.Println("s1= ", s1)
	s2 := Student{Person{"mile", 'w', 18}, 20, "jiangxi"}
	fmt.Println("s2= ", s2)
	//%+v显示的更加详细
	fmt.Printf("s2=%+v\n", s2)
	//指定成员初始化，没有初始化的自动赋值为0
	s3 := Student{id: 1}
	fmt.Printf("s=%+v\n", s3)

	s4 := Student{Person: Person{name: "mike"}, id: 1}
	fmt.Printf("s4=%+v\n", s4)

}
