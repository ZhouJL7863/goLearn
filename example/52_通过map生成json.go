package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["Company"] = "itcast"
	m["Subject"] = []string{"Go", "C++", "Python", "Test"}
	m["IsOk"] = true
	m["Price"] = 666.666
	//编码生成json
	//result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err= ", err)
		return
	}

	fmt.Println("result= ", string(result))
}
