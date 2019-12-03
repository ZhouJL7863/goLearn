package main

import "fmt"

func main() {
	a := 10
	str := "mike"
	func() {
		a = 666
		str = "go"
		fmt.Printf("a = %d,str = %s\n", a, str)
	}() //圆括号代表直接调用
}
