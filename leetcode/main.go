package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 0, 4, 5, 5}
	partition(arr, 3)
	fmt.Println(arr)
}

func partition(arr []int, pivot int) {
	//var small int = -1
	n := len(arr)
	var index int = 0
	for index != n {
		// if arr[index] < pivot {
		// 	small++
		// 	swap(arr, small, index)
		// 	fmt.Printf("swap %d and %d\n", small, index)
		// 	index++
		// } else if arr[index] == pivot {
		// 	index++
		if arr[index] <= pivot {
			index++
		} else {
			n--
			swap(arr, n, index)
			fmt.Printf("swap %d and %d\n", n, index)
		}
	}
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
