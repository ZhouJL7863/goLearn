package main

import "fmt"

type Hummaner interface { //子集
	sayhi()
}

type Personer interface { //超集
	Hummaner //匿名字段，继承了sayhi()
	sing(irc string)
}
type Student struct {
	name string
	id   int
}

func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s,%d] sayhi\n", tmp.name, tmp.id)

}
func (tmp *Student) sing(irc string) {
	fmt.Println("student 在唱着：", irc)
}
func main() {
	//超集可以转换为子集，反过来不可以
	var ipro Personer //超集
	ipro = &Student{"mile", 666}
	var i Hummaner //子集
	i = ipro
	i.sayhi()

}
