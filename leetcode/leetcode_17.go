package main

import "fmt"

//func leterCombinations(digits string) []string {
// var m = map[byte][]string{
// 	'2': []string{"a", "b", "c"},
// 	'3': []string{"d", "e", "f"},
// 	'4': []string{"g", "h", "i"},
// 	'5': []string{"j", "k", "l"},
// 	'6': []string{"m", "n", "o"},
// 	'7': []string{"p", "q", "r", "s"},
// 	'8': []string{"t", "u", "v"},
// 	'9': []string{"w", "x", "y", "z"},
// }
// if digits == "" {
// 	return nil
// }
// ret := []string{""} //表示的是含有一个字符串的数组
// //ret := []string{}表示的是一个空的字符串数组，
// fmt.Println(len(ret))
// for i := 0; i < len(digits); i++ {
// 	temp := []string{}
// 	for j := 0; j < len(ret); j++ {
// 		for k := 0; k < len(m[digits[i]]); k++ {
// 			temp = append(temp, ret[j]+m[digits[i]][k])
// 		}
// 	}
// 	ret = temp
// 	fmt.Println(ret)
// }
// return ret
// }
// func main() {
// digits := "23"
// ret := leterCombinations(digits)
// fmt.Println(ret)
// }

var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func leterCombinations(digits string) []string {
	ret := []string{}
	if len(digits) == 0 {
		return nil
	} else {
		temp := leterCombinations(digits[:len(digits)-1])
		if len(temp) == 0 {
			for k := 0; k < len(m[digits[len(digits)-1]]); k++ {
				ret = append(ret, m[digits[len(digits)-1]][k])
			}
		} else {
			for j := 0; j < len(temp); j++ {
				for k := 0; k < len(m[digits[len(digits)-1]]); k++ {
					ret = append(ret, temp[j]+m[digits[len(digits)-1]][k])
				}
			}
		}

	}
	return ret
}
func main() {
	digits := "234"
	ret := leterCombinations(digits)
	fmt.Println(ret)
	fmt.Println(len(ret))
}
