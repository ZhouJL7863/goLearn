package main

import "fmt"

func main() {
	x := 121
	s := isPalindrome(x)
	fmt.Println("s= ", s)
}
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	InverseNumber := 0
	for x > InverseNumber {
		InverseNumber = InverseNumber*10 + x%10
		x /= 10
	}
	return x == InverseNumber || x == InverseNumber/10
}
