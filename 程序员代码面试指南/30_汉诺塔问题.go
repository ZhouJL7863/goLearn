package main

import "fmt"

func process(arr []int, i, from, mid, to int) int {
	if i == -1 {
		return 0
	}
	if arr[i] != from && arr[i] != to {
		return -1
	}
	if arr[i] == from {
		return process(arr, i-1, from, to, mid)
	} else {
		rest := process(arr, i-1, mid, from, to)
		if rest == -1 {
			return -1
		}
		fmt.Println(i)
		return (1 << uint(i)) + rest
	}
}
func step1(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	return process(arr, len(arr)-1, 1, 2, 3)
}

func main() {
	arr := []int{3, 3}
	res := step2(arr)
	fmt.Println(res)
}

func step2(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	var from int = 1
	var to int = 3
	var mid int = 2
	i := len(arr) - 1
	var res int
	var tmp int
	for i >= 0 {
		if arr[i] != from && arr[i] != to {
			return -1
		}
		if arr[i] == to {
			res += 1 << uint(i)
			tmp = from
			from = mid

		} else {
			tmp = to
			to = mid
		}
		mid = tmp
		i--
	}
	return res
}
