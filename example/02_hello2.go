package main //必须有一个main包
import "fmt"

func main() {
	fmt.Println("hello go")
	s := []int{2, 4, 8}
	x := []int{4, 7}
	copy(s[1:], x)
	fmt.Println("s = ", s)
	ans := climbStairs(10)
	fmt.Println(ans)

	arr := multiple(20)
	fmt.Println(arr)
}
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func multiple(n int) int {
	if n == 1 {
		return 1
	}
	return n * multiple(n-1)
}
