package main

import "fmt"

func main() {
	a, b := 10, 20
	//通过一个函数交换a和b的内容
	//swap(&a, &b) //变量本身交换，值传递
	fmt.Printf("&a= %v,&b = %v\n", &a, &b)
	a, b = b, a

	fmt.Printf("main:a=%d,b=%d\n", a, b)
	fmt.Printf("&a= %v,&b = %v\n", &a, &b)
}
func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Printf("swap:a=%d,b=%d\n", *p1, *p2)

}
