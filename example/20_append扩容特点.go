package main

import "fmt"

func main() {
	//如果超过原来的容量，通常以2倍容量扩容
	s := make([]int, 0, 1) //容量为1,长度为0
	oldcap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if newcap := cap(s); oldcap < newcap {
			fmt.Printf("cap:%d====>%d\n", oldcap, newcap)
			oldcap = newcap
		}
	}
}
