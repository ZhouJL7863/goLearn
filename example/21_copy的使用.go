package main

import "fmt"

func main() {
	srcSlice := []int{1, 2, 3}
	dstSlice := []int{6, 6, 6, 6, 6}
	copy(dstSlice, srcSlice)
	fmt.Println("dst= ", dstSlice)

}
