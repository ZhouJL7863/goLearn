package main

import "fmt"

func main() {
	//声明定义同时赋值，叫初始化
	//全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("a= ", a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)
	//部分初始化，没有初始化的元素，自动赋值为0
	c := [5]int{1, 2, 3}
	fmt.Println("c= ", c)
	//arr := []int{},只是声明了，并没有初始化
	//arr := make([]int,len,cap),这样才会为数组分配内存空间
	d := []int{}
	d = append(d, 3)
	fmt.Println(d)

}
