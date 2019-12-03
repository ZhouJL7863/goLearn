package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "go"}
	//赋值，如果已经存在的key值，修改内容
	fmt.Println("m1 = ", m1)
	m1[1] = "c++"
	fmt.Println("m1= ", m1)

}
