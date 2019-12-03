package main

import (
	"fmt"
	"strings"
)

func main() {
	//"hellogo"中是否包含”hello"，包含返回True,不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))

	//Joins 组合
	s := []string{"abc", "mike", "go", "hello"}
	buf := strings.Join(s, " ")
	fmt.Println("buf= ", buf)

	//Index  查找子串的位置
	fmt.Println(strings.Index("abcdhell0", "hell0"))
	fmt.Println(strings.Index("abcdhell0", "hell0")) //不包含子串返回-1

	buf = strings.Repeat("go", 3)
	fmt.Println("buf= ", buf)

	//Split已指定的分隔符拆分
	buf = "hell0@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2= ", s2)

	//Trim去掉两头的字符
	buf = strings.Trim("   are you ok      ", " ") // 去掉空格
	fmt.Printf("buf= #%s#\n", buf)

	//去掉空格，把元素放入切片中
	s3 := strings.Fields("   are you ok?    ")
	//fmt.Println("s3= ",s3)
	for i, data := range s3 {
		fmt.Println(i, ", ", data)
	}
}
