package main

import "fmt"

//改变切片，底层数组也会改变
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:7]
	s1[2] = 444
	fmt.Println(arr)
	fmt.Println(s1)

	s2 := arr[2:7]
	s2[4] = 666
	fmt.Println(arr)
	fmt.Println(s2)
}
