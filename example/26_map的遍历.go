package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "youyou", 3: "c++"}
	for key, value := range m {
		fmt.Printf("%d======>%s\n", key, value)
	}
	//如何判断一个key值是否存在
	//第一个返回值为key所对应的value，第二个返回值是否存在的条件，存在🆗即为true
	value, ok := m[4]
	if ok == true {
		fmt.Println("m[1]= ", value)

	} else {
		fmt.Println("Key不存在")
	}
}

/*func twoSum(nums []int, target int) []int {
	m1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m1[nums[i]] = i
	}

	fmt.Println(m1)
	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		value, ok := m1[num]
		if ok == true && value != i {
			return []int{i, value}
		}
	}
	return []int{}
}
func main() {
	nums := []int{3, 2, 4}
	target := 6
	s := twoSum(nums, target)
	fmt.Println(s)
}
*/
