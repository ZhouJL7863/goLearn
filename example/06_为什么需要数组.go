package main

import "fmt"

func main() {
	//数组就是同一类型的集合
	var id [50]int
	//操作数组，通过下标，从0开始
	for i := 1; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}
}
