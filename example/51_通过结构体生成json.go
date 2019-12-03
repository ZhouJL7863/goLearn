package main

import (
	"encoding/json"
	"fmt"
)

// type IT struct {
// 	Company  string
// 	Subjects []string
// 	IsOk     bool
// 	Price    float64
// }
type IT struct {
	Company  string   `json："-"` //此字段不会输出到屏幕
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"bool"` //二次编码
	Price    float64  `json:"price"`
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"C++", "Go", "Python", "PHP", "peral"}, true, 666.666}

	//编码，根据内容生成JSON文本
	//buf, err := json.Marshaler(s)
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	fmt.Println("buf= ", string(buf))

}
