package main

import (
	"fmt"
)

//本题的思想是用双端队列来实现生成窗口最大值数组，其中双端队列的实现主要体现在getMaxWindow函数中两个if语句的实现。
//qmax是用来存放前面出现的最大值。
func getMaxWindow(arr []int, w int) []int {
	if len(arr) == 0 || w < 1 || len(arr) < w {
		return nil
	}
	qmax := []int{}
	res := make([]int, len(arr)-w+1)
	index := 0
	for i := 0; i < len(arr); i++ {
		for len(qmax) != 0 && arr[qmax[len(qmax)-1]] <= arr[i] {
			qmax = qmax[:len(qmax)-1]
		}
		qmax = append(qmax, i)
		if qmax[0] == i-w {
			qmax = qmax[1:]
		}
		if i >= w-1 {
			res[index] = arr[qmax[0]]
			index++
		}

	}
	return res
}

func main() {
	arr := []int{4, 3, 5, 4, 3, 3, 6, 7}
	var w int
	for {
		fmt.Println("请输入一个数字：")
		fmt.Scan(&w)
		res := getMaxWindow(arr, w)
		fmt.Println(res)
	}

}
