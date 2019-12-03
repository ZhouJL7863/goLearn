package main

import (
	"fmt"
)

func main() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型：%T\n", i)
	case int:
		fmt.Printf("x是int类型\n")
	case float64:
		fmt.Println("x是float64型")
	case func(int) float64:
		fmt.Println("x是func(int)型")
	case bool, string:
		fmt.Println("x是bool或string型")
	default:
		fmt.Println("未知型")
	}
	switch {
	case false:
		fmt.Println("1,case条件语句为false")
		fallthrough
	case true:
		fmt.Println("2,case条件语句为true")
		fallthrough
	case false:
		fmt.Println("3,case条件语句为false")
		fallthrough
	case true:
		fmt.Println("4,case条件语句为true")
	case false:
		fmt.Println("5,case条件语句为false")
		fallthrough
	default:
		fmt.Println("6默认为case")

	}
}

//这里有说明了go语言中虽然没用泛型的概念，但是提高了interface接口
//其中的在case中带有fallthrough关键字，程序将继续执行下一条case,且不会去判断
//下一个case的表示式是否为true.
