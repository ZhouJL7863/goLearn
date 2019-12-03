package main

import (
	"strconv"
)
import (
	"fmt"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	preString := countAndSay(n - 1)
	count := 1
	nowString := ""
	for i := 0; i <= len(preString)-1; i++ {
		if (i+1) < len(preString) && preString[i] == preString[i+1] { // 条件语句判断是有先后次序的
			count++
			continue
		} else {
			nowString += strconv.Itoa(count) + string(preString[i])
			count = 1
		}
	}
	return nowString
}
func main() {
	ans := countAndSay(6)
	fmt.Println("ans = ", ans)
}
