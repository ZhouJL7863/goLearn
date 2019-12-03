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
	s2 := Student{1, "mike", 'm', 18, "bj"}
	s3 := Student{2, "mike", 'm', 18, "bj"}
	fmt.Println("s1==s2", s1 == s2)
	fmt.Println("s1 == s3", s1 == s3)
	//同类型的2个结构体可以相互赋值
	var tmp Student
	tmp = s3
	fmt.Println("tmp= ", tmp)
}
