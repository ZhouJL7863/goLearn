package main

import "fmt"

func main() {
	a := 10
	str := "mike"

	//匿名函数，没有函数名字
	f1 := func() {
		fmt.Println("a= ", a)
		fmt.Println("str = ", str)
	}
	f1()

	func() {
		fmt.Printf("a = %d,str = %s\n", a, str)
	}() //后面的()代表调用此匿名函数
	//带参数的匿名函数
	f3 := func(i, j int) {
		fmt.Printf("i=%d, j=%d\n", i, j)
	}(10, 20)
	//匿名函数有返回值
	func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)
}
