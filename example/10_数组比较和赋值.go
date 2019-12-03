package main

import "fmt"

func main() {
	//支持比较，只支持 == 或 ！=，
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}
	fmt.Println("a==b", a == b)
	fmt.Println("a==c", a == c)
	var d [5]int
	d = a
	fmt.Println("d = ", d)
}
