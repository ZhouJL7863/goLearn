package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max]取下标从low开始的元素，
	s1 := array[:]
	fmt.Printf("s1= ", s1)
	fmt.Println("len=%d,cap=%d\n", len(s1), cap(s1))
	//操作某个元素，和数组元素的操作方式一样
	data := array[1]
	fmt.Println("data= ", data)

	s2 := array[3:6:7]
	fmt.Println("s1= ", s1)
	fmt.Printf("len=%d,cap=%d\n", len(s2), cap(s2))

}
