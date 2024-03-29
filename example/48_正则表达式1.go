package main

import (
	"fmt"
	"regexp"
)

func main() {
	//1)解释规则

	buf := "abc azc a7c aac 888 a9c 5ac"
	//reg1 := regexp.MustCompile("a.c")

	reg1 := regexp.MustCompile("a[0-9]c")
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	//2）根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1= ", result1)
}
