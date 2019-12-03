package main

import "fmt"

func main() {
	//变量，程序运行期间，可以改变的量
	// 1.声明格式， var变量名 类型，变量声明了，必须使用
	var a int
	fmt.Println(a)
	fmt.Println("a = ", a)
	//变量的初始化，声明变量时，同时赋值
	var b int = 10
	b = 20
	fmt.Println("b = ", b)
	//自动推导类型，必须初始化，通过初始化的值确定类型
	c := 30
	fmt.Print("c type is %T\n", c)
}
