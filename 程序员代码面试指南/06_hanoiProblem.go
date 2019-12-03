package main

import (
	"fmt"
	"strconv"
)

func hanoiProblem(num int, left, mid, right string) int {
	if num < 1 {
		return 0
	}
	return process(num, left, mid, right, left, right)
}

func process(num int, left, mid, right, from, to string) int {
	if num == 1 {
		if from == mid || to == mid {
			fmt.Println("Move 1 from " + from + " to " + to)
			return 1
		} else {
			fmt.Println("Move 1 from " + from + " to " + mid)
			fmt.Println("Move 1 from " + mid + " to " + to)
			return 2
		}
	}
	if from == mid || to == mid {
		var another string
		if from == left || to == left {
			another = right
		} else {
			another = left
		}
		part1 := process(num-1, left, mid, right, from, another)
		part2 := 1
		fmt.Println("Move " + strconv.Itoa(num) + " from " + from + " to " + to)
		part3 := process(num-1, left, mid, right, another, to)
		return part1 + part2 + part3
	} else {
		part1 := process(num-1, left, mid, right, from, to)
		part2 := 1
		fmt.Println("Move " + strconv.Itoa(num) + " from " + from + " to " + mid)
		part3 := process(num-1, left, mid, right, to, from)
		part4 := 1
		fmt.Println("Move " + strconv.Itoa(num) + " from " + mid + " to " + to)
		part5 := process(num-1, left, mid, right, from, to)
		return part1 + part2 + part3 + part4 + part5
	}

}
func main() {
	left := "left"
	right := "right"
	mid := "mid"
	ret := hanoiProblem(2, left, mid, right)
	fmt.Printf("It will move %d steps\n", ret)
}
