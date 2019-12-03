package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 agsdg 1.23 7.  8 gajsdj gjsd j 3441.534"
	//解释正则表达式
	reg := regexp.MustCompile(`[0-9]+\.[0-9]+`)
	if reg == nil {
		fmt.Println("MustCompile Err")
		return
	}
	//提取关键信息
	//result := reg.FindAllStrings(buf,-1)
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result", result)
}
