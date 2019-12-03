package main

import "fmt"

type Hummaner interface {
	sayhi()
}

type Personer interface {
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
	var i Personer
	s := &Student{"mile", 666}
	i = s
	i.sayhi()
	i.sing("学生歌")
}
