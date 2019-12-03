package main

import "fmt"
import "math/rand"
import "time"

func main() {
	var a [10]int
	rand.Seed(time.Now().UnixNano()) //获取当前时间的种子
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(100) //100以内的随机数
	}
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}

		}
	}
	fmt.Printf("a = ", a)
}
