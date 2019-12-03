package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	modify(&a) //数组传递过去
	fmt.Println("main:a= ", a)
}

//数组做函数参数，它是值传递
//实参数组的每个元素个形参数组拷贝一份
func modify(p *[5]int) {
	(*p)[0] = 666
	fmt.Println("modify *a= ", *p)

}
