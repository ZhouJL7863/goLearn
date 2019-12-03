//滑动窗口的使用，还是出了点问题
//字符串的切片还是字符串，字符串的索引是字节
package main

import (
	"fmt"
	//"strings"
	//"strings"
)

func main() {
	s := "abba"
	ans := lengthOfLongestSubstring(s)
	fmt.Println("ans = ", ans)
	for _, value := range s {
		fmt.Println("value", string(value))
		fmt.Printf("%T\n", value)
	}
}

//这是使用MAP的用法，下面使用strings的性质
func lengthOfLongestSubstring(s string) int {
	arr := make(map[string]int)
	high := 0
	low := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		high = i
		if _, ok := arr[string(s[i])]; ok {
			low = Max(arr[string(s[i])]+1, low)
			//low = arr[string(s[i])] + 1
		}
		arr[string(s[i])] = i
		ans = Max(high-low+1, ans)
		fmt.Println(ans, low)
	}
	return ans
}
func Max(max, min int) int {
	if max > min {
		return max
	}
	return min
}
