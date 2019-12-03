package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	if target == 0 {
		temp := []int{0}
		ret = append(ret, temp)
		return ret
	}
	if target < candidates[0] {
		return [][]int{}
	}

	for i := 0; i < len(candidates); i++ {
		arr := combinationSum(candidates, target-candidates[i])
		fmt.Println(ret)
		n := len(arr)
		if n == 0 {
			return [][]int{}
		} else {
			for j := 0; j < n; j++ {
				arr[j] = append(arr[j], candidates[i])
				temp := make([]int, len(arr[j]))
				copy(temp, arr[j])
				ret = append(ret, temp)
				fmt.Println(ret)
			}
		}

	}
	return ret
}
func main() {
	candidates := []int{2, 3, 6, 7}
	target := 2
	ret := combinationSum(candidates, target)
	fmt.Println(ret)
}
