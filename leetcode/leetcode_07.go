package main

import "fmt"

func main() {
	x := -2143847412
	result := inverse(x)

	fmt.Println("result= ", result)

}
func inverse(x int) int {
	max := 1
	min := 1
	for i := 0; i < 31; i++ {
		max *= 2
		min *= 2
	}
	max = max - 1
	min = -min
	ans := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if (ans > max/10) || (ans == max/10 && pop > 7) {
			return 0
		}
		if (ans < min/10) || (ans == min/10 && pop < (-8)) {
			return 0
		}

		ans = ans*10 + pop
	}
	return ans
}
