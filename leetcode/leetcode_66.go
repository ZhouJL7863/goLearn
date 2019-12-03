package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4} //这是数组
	res := plusOne(s)
	fmt.Println("res= ", res)
	x := printArray()
	fmt.Println("x= ", x)

}
func plusOne(s []int) []int {
	//n := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < 9 {
			s[i]++
			return s
		}
		s[i] = 0
	}
	//整体产生了进位，数组的长度需要加1
	res := make([]int, len(s)+1)
	//不知道为什么直接进行赋值就不行了
	//var res [4]int,这样的报错是因为返回值为[]int，而数组的长度是固定的，但是可以用切片，不知道是什么原因。
	//数组是[]int里面的数字是固定的，因此不能用不能表示的是同一返回值，而用切片就可以表示长度和容量。
	res[0] = 1
	return res
}
func printArray() []int {
	s := []int{1, 2, 3, 4}
	res := make([]int, 4)
	copy(res, s) //，参数必须同时是
	//for i := 0; i < len(s); i++ {
	//res[i] = s[i]
	//}
	return res
}
func main01() {
	s := []int{1, 2, 3, 4}
	p := []int{2, 3, 4, 5, 6, 7}
	copy(p, s)

	fmt.Println("p= ", p)
}
