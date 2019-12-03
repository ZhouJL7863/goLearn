package main

import "fmt"

func main() {
	result := myAtoi("123555")
	fmt.Println("result= ", result)
}
func myAtoi(str string) int {
	const (
		int32Max = 1<<31 - 1
		int32Min = -1 << 31
	)

	n := len(str)
	var i, j int
	neg := false //用来区分正负数的
	for i = 0; i < n; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			break
		} else if str[i] == '+' {
			i++
			break
		} else if str[i] == '-' {
			neg = true
			i++
			break
		} else if str[i] != ' ' {
			return 0
		}
	}
	for j = i; j < n; j++ {
		if str[j] < '0' || str[j] > '9' {
			break
		}
	}
	ret := 0
	for k := i; k < j; k++ {
		cur := int(str[k] - '0')
		ret = ret*10 + cur
		if !neg {
			if ret > int32Max {
				return int32Max
				break
			}
		} else {
			if (-ret) < int32Min {
				return int32Min
				break
			}
		}
	}
	if neg == false {
		return ret
	} else {
		return -ret
	}
}
