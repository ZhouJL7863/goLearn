package main

import (
	"encoding/json"
	"fmt"
)

// type IT struct {
// 	company  string   `json："-"` //此字段不会输出到屏幕
// 	subjects []string `json:"subjects"`
// 	isok     bool     `json:"isok"` //二次编码
// 	price    float64  `json:"price"`
// }

func main() {
	jsonBuf := `
	{
		"company":"itcast",
		"subjects": [
			"Go",
			"C++",
			"Python",
			"Test"
		],
		"isok":true,
		"price":666.666
	
	}`
	m := make(map[string]interface{}, 4)
	//var tmp IT //定义一个结构体变量
	//err := json.Unmarshal([]byte(jsonBuf), &tmp)  //第二个参数要地址传递
	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	//fmt.Println("tmp= ", tmp)
	fmt.Printf("m = %+v\n", m)
	//类型断言
	//var str string
	for key, value := range m {
		fmt.Printf("%v ========>%v\n", key, value)
		switch data := value.(type) {
		case string:
			//str = data
			fmt.Printf("map[%s]的值类型为string,value= %s\n", key, data)
		case bool:
			fmt.Printf("map[%s]的值类型为bool,value= %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64,value= %f\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为float64,value= %v\n", key, data)
		}
	}
}

//文件的分类：设备文件：屏幕(标准输出设备）fmt.Println()往标准输出设备写内容
//键盘（标准输入设备）fmt.Scan() 从标准输入设备读取内容
//磁盘文件，放在存储设备上的文件
//1）文本文件，以记事本打开，能看到内容（不是乱码） 2）.二进制文件，以记事本打开，能看到内容（是乱码）
//为什么需要文件？内存掉电丢失，程序结束，内存中的内容消失，文件放磁盘，程序结束，文件还是存在。
