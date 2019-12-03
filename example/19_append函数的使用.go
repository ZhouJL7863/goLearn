package main

import "fmt"

func main() {
	a1 := []int{}
	fmt.Printf("len=%d,cap=%d\n", len(a1), cap(a1))
	fmt.Println("s1= ", a1)
	//在原切片的末尾添加元素
	a1 = append(a1, 1)
	a1 = append(a1, 2)
	a1 = append(a1, 3)
	fmt.Printf("len=%d,cap=%d\n", len(a1), cap(a1))

	s2 := []int{1, 2, 3}
	fmt.Println("s2= ", s2)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	fmt.Println("s2= ", s2)
	fmt.Printf("len=%d,cap=%d\n", len(s2), cap(s2))

	nums := []int{1, 2, 3}
	fmt.Println("****************")
	Add(nums)
	fmt.Println(nums)
	fmt.Printf("%p", &nums)

}

//切片做函数参数，应用类型：顺序可以改变，但是扩容就不行
func Add(nums []int) {
	nums[1], nums[0] = nums[0], nums[1]
	nums = append(nums, 4)
	fmt.Printf("%p\n", &nums)
}
