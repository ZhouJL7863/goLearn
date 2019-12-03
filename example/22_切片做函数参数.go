package main

import "fmt"
import "math/rand"
import "time"

func main() {
	var n int = 10
	//创建一个切片，len为n
	s := make([]int, n)
	InitData(s) //初始化数组
	fmt.Println("排序前: ", s)
	Bubblesort(s) //冒泡排序
	fmt.Println("排序后:", s)
}
func Bubblesort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
func InitData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100) //100以内的随机数
	}
}
