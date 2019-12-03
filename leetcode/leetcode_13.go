package main

import "fmt"

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
func romanToInt(s string) int {
	arr := make(map[string]int)
	arr = map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	Num := 0
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %s\n", i, string(s[i]))
		fmt.Printf("%T\n", s[i])
	}
	for i := 0; i < len(s); i++ {
		y := Min(i+2, len(s))
		if value, ok := arr[s[i:y]]; ok == true {
			i++
			Num += value
		} else if value, ok := arr[string(s[i])]; ok == true {
			Num += value
		}

	}

	return Num
}
func Min(max, min int) int {
	if max > min {
		return min
	}
	return max
}

//记住字符串的输出是整型
//只有用string才能转换为字符，同时切片是字符串类型
