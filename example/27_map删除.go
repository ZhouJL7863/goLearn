package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "youyou"}
	fmt.Println("m= ", m)
	delete(m, 1)
	fmt.Println("m= ", m)
}
