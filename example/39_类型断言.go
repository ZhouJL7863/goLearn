package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1 //int
	i[1] = "hello go"
	i[2] = Student{"mike", 666}
	//类型查询，类型断言。
	//第一个是下标，第二个是返回下标的对应的值，data分别是i[0],i[1],i[2]
	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int，内容为%s\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string，内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d],类型为Student,内容为name= %s,id= %d\n", index, value.name, value.id)
		}
	}
}
