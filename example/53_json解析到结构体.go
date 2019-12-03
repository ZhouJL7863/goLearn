package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json："-"` //此字段不会输出到屏幕
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok"` //二次编码
	Price    float64  `json:"price"`
}

func main() {
	jsonBuf := `
	{
		"company": "itcast",
		"subjects": [
			"Go",
			"C++",
			"Python",
			"Test"
		],
		"isok": true,
		"price": 666.666
	
	}`
	var tmp IT //定义一个结构体变量
	//err := json.Unmarshal([]byte(jsonBuf), &tmp)  //第二个参数要地址传递
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	fmt.Println("tmp= ", tmp)
}
