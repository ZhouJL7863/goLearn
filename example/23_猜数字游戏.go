package main

import "fmt"
import "math/rand"
import "time"

func main() {
	var randNum int
	randNum = GetNum()
	NumSlice := make([]int, 4)
	GetSlice(NumSlice, randNum)
	//fmt.Println("NumSlice= ", NumSlice)

	for {

		n := 0
		Num := OnGame()
		KyeSlice := make([]int, 4)
		GetSlice(KyeSlice, Num)
		for i := 0; i < len(KyeSlice); i++ {
			if NumSlice[i] > KyeSlice[i] {
				fmt.Printf("你输入的第%d数字小了\n", i+1)
			} else if NumSlice[i] < KyeSlice[i] {
				fmt.Printf("你输入的第%d数字大了\n", i+1)
			} else {
				fmt.Printf("你输入的第%d数字正确\n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("你输入的数字全部对了")
			break
		}
	}
}
func GetNum() (randNum int) {
	rand.Seed(time.Now().UnixNano())
	for {
		randNum = rand.Intn(10000)
		if randNum > 999 && randNum < 10000 {
			break
		}
	}
	return
}
func GetSlice(s []int, num int) {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}
func OnGame() (num int) {
	for {
		fmt.Println("请输入一个数字：")
		fmt.Scan(&num)
		if num > 999 && num < 10000 {
			break
		}
		fmt.Println("你输入的数字不符合要求，请重新输入：")
	}
	return num
}
