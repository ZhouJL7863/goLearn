package main

import (
	"fmt"
	//"strings"
)

// func main() {
// s := "我是中国人"
// fmt.Println("utf-8 string: ", s, len(s))
// for i := 0; i < len(s); i++ {
// 	fmt.Println("s = ", s[i])
// }
// for x, y := range s {
// 	fmt.Printf("%d = %c\n", x, y)
// }
//转换为bytes，以便于修改
// bs := []byte(s)
// bs[0] = 'A'
// for i := 0; i < len(s); i++ {
// 	fmt.Printf("%02x\n", bs[i])
// }
// fmt.Print(string(bs))

// u := []rune(s)
// fmt.Println("unicode len= ", len(u))
// for i := 0; i < len(u); i++ {
// 	fmt.Printf("%04x\n", u[i])
// }
// fmt.Println(string(u))
//s := []int{1, 2, 3, 4}
//fmt.Print(s[0:1])
// X := []string{"a", "b", "c"}
// for i := 0; i < 3; i++ {
// 	y := X[i]
// 	fmt.Printf("%T\n", y)
// 	fmt.Println("s=", X[i])
// }
// x := 6
// y := 11
// fmt.Println(x ^ y)
// s := []int{10, 3, 3, 3}
// ans := singleNumber(s)
// fmt.Println("s = ", ans)

// }
// func singleNumber(s []int) int {
// var ones, twos int = 0, 0
// for i := 0; i < len(s); i++ {
// 	ones = (s[i] ^ ones) &^ (twos)
// 	twos = (s[i] ^ twos) &^ (ones)
// }
// return ones
// }
func backTrack(n, first int, nums []int, res *[][]int) {
	if first == n {
		temp := make([]int, n)
		copy(temp, nums)
		fmt.Println(temp)
		*res = append(*res, temp)
	}
	for i := first; i < n; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backTrack(n, first+1, nums, res)
		nums[first], nums[i] = nums[i], nums[first]
	}
}
func permte(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	first := 0
	backTrack(n, first, nums, &res)
	return res
}
func main() {
	nums := []int{1, 2, 3}
	res := permte(nums)
	fmt.Println(res)
}
