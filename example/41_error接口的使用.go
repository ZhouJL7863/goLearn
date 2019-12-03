package main

import "fmt"
import "errors"

func main() {
	//fmt.Errorf("%s", "this is normol error")
	err1 := fmt.Errorf("%s", "this is normol error")
	fmt.Println("err1= ", err1)

	err2 := errors.New("this is normal err2")
	fmt.Println("err2= ", err2)

}
