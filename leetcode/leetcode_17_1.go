package main

import "fmt"
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
func backTrack(combination,next_digits string) {
	if  len(next_digits) == 0 {
		output.add(combination)
	}else {
		digit := next_digits[0]
		letter := map[digit]
		for i:=0;i<len(letter);i++{
			backTrack(combination+letter[i],next_digits[1:])
		}
	}
}
func letterCombinations(digits string)(output []string){
	if len(digits) != 0 {
		backTrack("",digits)
		return output
	}
}